package service

import (
	"backend/internal/dao"
	"backend/internal/dto"
	"backend/internal/model"
	"backend/pkg/agent"

	// 由于缺少 backend/pkg/agent 的元数据，暂时注释掉该导入
	// "backend/pkg/agent"
	"context"
	"encoding/json"
	"fmt"
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
}

func NewAgentService(client *react.Agent, fileParser *openai.Client, db *gorm.DB, redis *redis.Client) *AgentService {
	return &AgentService{
		Client:          client,
		FileParser:      fileParser,
		ConversationDao: dao.NewConversationDao(db, redis),
	}
}

func (s *AgentService) Chat(req *dto.ChatRequest, message []*schema.Message, c *gin.Context) (string, error) {
	var fullcontent string
	resp, err := s.Client.Stream(context.Background(), message)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	messchan := make(chan string)
	defer func() {
		// 保存到数据库
		s.ConversationDao.AddMessageToMysql(&model.Message{
			ConversationId: req.ConversationId,
			UserId:         req.UserID,
			Role:           schema.Assistant,
			Content:        fullcontent,
		})
		message = append(message, &schema.Message{
			Role:    schema.Assistant,
			Content: fullcontent,
		})
		close(messchan)

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
				break
			} else {
				return fullcontent, err
			}
		}
		messchan <- mess.Content
		fullcontent += mess.Content
		fmt.Println(mess.Content)
	}
	return fullcontent, nil
}

func Stream(c *gin.Context, mess chan string) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	for {
		select {
		case msg := <-mess:
			jsonMsg := struct {
				Data string `json:"data"`
			}{msg}
			jsonm, _ := json.Marshal(jsonMsg)
			fmt.Fprintf(c.Writer, "data: %s\n\n", string(jsonm))
			c.Writer.Flush()
		case <-c.Done():
			return
		}
	}
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

func (s *AgentService) ParseFile(fileIDs map[string]string, client *openai.Client, userMessage string, c *gin.Context) (schema.Message, error) {
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
		return schema.Message{}, err
	}
	// for _, choice := range resp.Choices {
	// 	parsedFiles[choice.Message.Content] = choice.Message.Content
	// }
	// fmt.Println(resp.Choices[0].Message.Content)
	// return schema.Message{
	// 	Role:    schema.Assistant,
	// 	Content: resp.Choices[0].Message.Content,
	// }, nil
	messchan := make(chan string)
	var content string
	go Stream(c, messchan)
	for {
		select {
		case <-c.Request.Context().Done():
			return schema.Message{}, nil
		default:
		}
		mess, err := resp.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return schema.Message{}, err
			}
		}
		content += mess.Choices[0].Delta.Content
		messchan <- mess.Choices[0].Delta.Content
	}
	return schema.Message{
		Role:    schema.Assistant,
		Content: content,
	}, nil
}
