package service

import (
	"backend/internal/dao"
	"backend/internal/dto"
	"backend/internal/model"
	"backend/pkg/agent"
	"bytes"
	"context"
	"fmt"
	eino "github.com/cloudwego/eino-ext/components/model/openai"
	"html/template"
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
	tp1, err := template.New("prompt").Parse(agent.ParseUserPicturePrompt)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var b bytes.Buffer
	job_position, err := dao.GetJobAndPositon()
	data := struct {
		JobPosition []dto.GetPositionByJobRes `json:"job_position"`
	}{
		JobPosition: job_position,
	}
	err = tp1.Execute(&b, data)
	if err != nil {
		return nil, err
	}
	system := b.String()
	fmt.Println(system)
	requestMessage := make([]openai.ChatCompletionMessage, 0)
	requestMessage = append(requestMessage, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: system,
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
		Stream:   false,
	}
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var content string
	content = resp.Choices[0].Message.Content
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

func (s *AgentService) GetPositionDifference(message []*schema.Message, client *eino.ChatModel, c *gin.Context) error {
	ctx := context.Background()
	ch := make(chan string, 10)
	go Stream(c, ch)
	stream, err := client.Stream(ctx, message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for {
		message, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				close(ch)
				break
			}
			return err
		}
		ch <- message.Content
		fmt.Print(message.Content)
	}
	return nil
}

func (h *AgentService) ParseResume(c *gin.Context) {
	var user dto.UserPictureRes
	var position model.Position
	user, _ = dao.GetUserPictureByUserId(c.GetUint("id"))
	position, _ = dao.GetJobPicture("C/C++ 嵌入式系统方向")
	mp := map[string]interface{}{
		"UserProfile": fmt.Sprintf("%v", user),
		"TargetJob":   fmt.Sprintf("%v", position),
	}
	ch := make(chan string, 10)
	go Stream(c, ch)
	err := agent.GenerateResumeInsight(mp, ch, h.ChatModel)
	if err != nil {
		return
	}
}

func (h *AgentService) GetResumeRadar(c *gin.Context, userId uint, job string) (dto.ResumeRadarRes, error) {
	var user dto.UserPictureRes
	var position model.Position
	user, _ = dao.GetUserPictureByUserId(userId)
	position, _ = dao.GetJobPicture(job)
	mp := map[string]interface{}{
		"candidate_profile": fmt.Sprintf("%v", user),
		"job_profile":       fmt.Sprintf("%v", position),
	}
	radar, err := agent.GenerateResumeRadar(mp, h.ChatModel)
	if err != nil {
		return dto.ResumeRadarRes{}, err
	}
	return radar, nil
}
