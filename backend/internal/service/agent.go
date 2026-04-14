package service

import (
	"backend/internal/dao"
	"backend/internal/dto"
	"backend/internal/model"
	"backend/pkg/agent"
	"context"
	"fmt"
	eino "github.com/cloudwego/eino-ext/components/model/openai"
	"io"
	"os"
	"path/filepath"

	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

type AgentRequest struct {
	UserMessage string            `json:"user_message" description:"用户消息"`
	FileIDs     map[string]string `json:"file_ids" description:"文件ID"`
}

type AgentService struct {
	Client          *react.Agent
	FileParser      *openai.Client
	ConversationDao *dao.ConversationDao
	ChatModel       *eino.ChatModel
}

func NewAgentService(client *react.Agent, fileParser *openai.Client, db *gorm.DB, redis *redis.Client, chatModel *eino.ChatModel) *AgentService {
	return &AgentService{
		Client:          client,
		FileParser:      fileParser,
		ConversationDao: dao.NewConversationDao(db, redis),
		ChatModel:       chatModel,
	}
}

func (s *AgentService) ChatV2(req *dto.ChatRequest, c *gin.Context) error {
	ch := make(chan string, 10)
	career := agent.NewCareer(req.UserID, req.Message, s.ConversationDao)
	ctx := context.Background()
	go Stream(c, ch)
	err := career.Graph(ctx, career, s.ChatModel, ch)
	if err != nil {
		return err
	}
	return nil
}

func (s *AgentService) Chat(req *dto.ChatRequest, message []*schema.Message, c *gin.Context) (string, error) {
	var fullcontent string
	for _, message := range message {
		fmt.Println(message.Role)
	}
	resp, err := s.Client.Stream(context.Background(), message)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	messchan := make(chan string)
	defer func() {
		// 保存到数据库
		s.ConversationDao.AddMessageToMysql(&model.Message{
			UserId:  req.UserID,
			Role:    schema.Assistant,
			Content: fullcontent,
		})
		// 保存到Redis
		s.ConversationDao.AddMessageToRedis(req.UserID, &model.Message{
			UserId:  req.UserID,
			Role:    schema.Assistant,
			Content: fullcontent,
		})
	}()
	go Stream(c, messchan)
	for {
		select {
		case <-c.Request.Context().Done():
			return fullcontent, nil
		default:
		}
		mess, err := resp.Recv()
		if err != nil {
			if err == io.EOF {
				close(messchan)
				break
			} else {
				return fullcontent, err
			}
		}
		messchan <- mess.Content
		fullcontent += mess.Content
	}
	return fullcontent, nil
}

func Stream(c *gin.Context, mess chan string) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Stream(func(w io.Writer) bool {
		select {
		case msg, ok := <-mess:
			if !ok {
				return false
			}
			jsonMsg := struct {
				Data string `json:"content"`
			}{msg}
			c.SSEvent("json", jsonMsg)
			return true
		case <-c.Request.Context().Done():
			return false
		}
	})
}

func (s *AgentService) UploadFile(file_path []string, client *openai.Client) (map[string]string, error) {
	// uploadReq := openai.FileRequest{
	// 	FilePath: file_path[0],
	// 	Purpose:  "file-extract",
	// 	FileName: filepath.Base(file_path[0]),
	// }
	// file,err:=client.CreateFile(context.Background(),uploadReq)
	// if err!= nil{
	// 	fmt.Println(err)
	// 	return err
	// }
	// err=os.Remove(file_path[0])
	// if err!= nil{
	// 	fmt.Println(err)
	// 	return err
	// }
	var fileIDs map[string]string = make(map[string]string)
	for _, path := range file_path {
		uploadReq := openai.FileRequest{
			FilePath: path,
			Purpose:  "file-extract",
			FileName: filepath.Base(path),
		}
		file, err := client.CreateFile(context.Background(), uploadReq)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		fmt.Println("上传成功")
		err = os.Remove(path)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		fileIDs[filepath.Base(path)] = file.ID
	}
	return fileIDs, nil
}

func (s *AgentService) ParseFile(fileIDs map[string]string, client *openai.Client, userMessage string, c *gin.Context) ([]*schema.Message, error) {
	// var parsedFiles map[string]string = make(map[string]string)
	requestMessage := make([]openai.ChatCompletionMessage, 0)
	requestMessage = append(requestMessage, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: agent.FilePrompt,
	})
	for _, fileID := range fileIDs {
		requestMessage = append(requestMessage, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: fmt.Sprintf("fileid://%s", fileID),
		})
	}
	requestMessage = append(requestMessage, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: userMessage,
	})
	req := openai.ChatCompletionRequest{
		Model:    os.Getenv("Qwen_longModel"),
		Messages: requestMessage,
		Stream:   true,
	}
	resp, err := client.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// for _, choice := range resp.Choices {
	// 	parsedFiles[choice.Message.Content] = choice.Message.Content
	// }
	// fmt.Println(resp.Choices[0].Message.Content)
	// return schema.Message{
	// 	Role:    schema.Assistant,
	// 	Content: resp.Choices[0].Message.Content,
	// }, nil
	var content string
	for {
		select {
		case <-c.Request.Context().Done():
			return []*schema.Message{
				{
					Role:    schema.User,
					Content: userMessage,
				},
				{
					Role:    schema.Assistant,
					Content: content,
				},
			}, nil
		default:
		}
		mess, err := resp.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		content += mess.Choices[0].Delta.Content
	}
	return []*schema.Message{
		{
			Role:    schema.User,
			Content: userMessage,
		},
		{
			Role:    schema.Assistant,
			Content: content,
		},
	}, nil
}
