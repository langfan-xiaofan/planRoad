package handler

import (
	"backend/internal/dao"
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/service"
	"encoding/json"
	"fmt"
	eino "github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

type AgentHandler struct {
	AgentService *service.AgentService
}

func NewAgentHandler(client *react.Agent, fileParser *openai.Client, db *gorm.DB, redis *redis.Client, chatModel *eino.ChatModel) *AgentHandler {
	return &AgentHandler{
		AgentService: service.NewAgentService(client, fileParser, db, redis, chatModel),
	}
}
func (h *AgentHandler) Chat(c *gin.Context) {
	var req dto.ChatRequest
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req.UserID = uint(c.GetInt("userid"))
	//convId, err := strconv.Atoi(form.Value["conversation_id"][0])
	if err != nil {
		c.JSON(400, gin.H{"error": "非法的对话ID"})
		return
	}
	fmt.Println(req)
	// 处理文件上传
	file := form.File["file"]
	filePaths := make([]string, 0)
	for _, f := range file {
		if err := c.SaveUploadedFile(f, fmt.Sprintf("./uploads/%s", f.Filename)); err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		filePaths = append(filePaths, fmt.Sprintf("./uploads/%s", f.Filename))
	}
	// 若没有消息，默认分析文件
	req.Message = form.Value["message"][0]
	if req.Message == "" {
		req.Message = "请分析这个文件"
	}
	// 上传文件并获取文件ID
	fileIDs, err := h.AgentService.UploadFile(filePaths, h.AgentService.FileParser)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// 解析文件返回文件内容
	parsedFiles, err := h.AgentService.ParseFile(fileIDs, h.AgentService.FileParser, req.Message, c)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// 添加解析后的文件内容消息到Redis和数据库
	for _, message := range parsedFiles {
		h.AgentService.ConversationDao.AddMessageToRedis(req.UserID, &model.Message{
			UserId:  req.UserID,
			Role:    message.Role,
			Content: message.Content,
		})
		h.AgentService.ConversationDao.AddMessageToMysql(&model.Message{
			UserId:  req.UserID,
			Role:    message.Role,
			Content: message.Content,
		})
	}
	// 从Redis获取历史消息，若没有则从数据库获取,若数据库也没有则创建新对话
	var messages []*schema.Message
	var historyMessages []*schema.Message
	messages = append(messages, &schema.Message{
		Role:    schema.System,
		Content: "你是职业规划助手，你的任务是根据用户上传的文件，分析文件内容，为用户提供职业规划建议",
	})
	historyMessages, err = h.AgentService.ConversationDao.GetMessagesFromRedis(req.UserID)
	if err != nil {
		historyMessages, err = h.AgentService.ConversationDao.GetMessagesFromMysql(req.UserID)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
	// 合并历史消息和解析后的文件内容消息
	for _, message := range historyMessages {
		messages = append(messages, message)
	}
	// 若对话消息超过20条，总结一下以上内容
	defer func() {
		if len(messages) >= 20 {
			newMessages := make([]*schema.Message, 0)
			newMessages = append(newMessages, &schema.Message{
				Role:    schema.User,
				Content: "请总结一下以上内容",
			})
			messages = append(messages, &schema.Message{
				Role:    schema.User,
				Content: "请总结一下以上内容",
			})
			newMess, err := h.AgentService.Chat(&req, messages, c)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			newMessages = append(newMessages, &schema.Message{
				Role:    schema.Assistant,
				Content: newMess,
			})
			for _, message := range newMessages {
				h.AgentService.ConversationDao.AddMessageToRedis(req.UserID, &model.Message{
					UserId:  req.UserID,
					Role:    message.Role,
					Content: message.Content,
				})
			}
		}
	}()
	// 添加用户消息到对话，同时添加到Redis和数据库
	messages = append(messages, &schema.Message{
		Role:    schema.User,
		Content: req.Message,
	})
	h.AgentService.ConversationDao.AddMessageToMysql(&model.Message{
		UserId:  req.UserID,
		Role:    schema.User,
		Content: req.Message,
	})
	h.AgentService.ConversationDao.AddMessageToRedis(req.UserID, &model.Message{
		UserId:  req.UserID,
		Role:    schema.User,
		Content: req.Message,
	})
	// 调用Chat方法获取回复
	if _, err = h.AgentService.Chat(&req, messages, c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (h *AgentHandler) Parse(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	file := form.File["file"]
	filePaths := make([]string, 0)
	for _, f := range file {
		if err := c.SaveUploadedFile(f, fmt.Sprintf("./uploads/%s", f.Filename)); err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		filePaths = append(filePaths, fmt.Sprintf("./uploads/%s", f.Filename))
	}
	userMessage := form.Value["personalIntro"][0]
	if userMessage == "" {
		userMessage = "请分析这个文件"
	}
	fileIDs, err := h.AgentService.UploadFile(filePaths, h.AgentService.FileParser)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	parsedFiles, err := h.AgentService.ParseFile(fileIDs, h.AgentService.FileParser, userMessage, c)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("分析出来的画像", parsedFiles[1].Content)
	var position dto.UserPictureReq
	err = json.Unmarshal([]byte(parsedFiles[1].Content), &position)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	position.UserId = c.GetUint("id")
	err = dao.CreatUserPicture(position)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{"parsedFiles": position})
}

func (h *AgentHandler) ChatV2(c *gin.Context) {
	var req dto.ChatRequest
	form, err := c.MultipartForm()
	req.UserID = uint(c.GetInt("user_id"))
	req.Message = form.Value["message"][0]
	fmt.Println("请求：", req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	err = h.AgentService.ChatV2(&req, c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true})
}

func (h *AgentHandler) GetPositionDeference(c *gin.Context) {
	var df dto.GetPositionDifferenceRequest
	if err := c.ShouldBindJSON(&df); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	messages := make([]*schema.Message, 0)
	messages = append(messages, &schema.Message{
		Role:    schema.System,
		Content: "你是一个专业的职位对比分析师，你将根据两个职业画像分析从第一个职业转到第二个职业所需要学习的技能，以及一些建议，要求精简",
	})
	var po1, po2 model.Position
	po1, err := dao.GetJobPicture(df.From)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	po2, err = dao.GetJobPicture(df.To)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	messages = append(messages, &schema.Message{
		Role:    schema.User,
		Content: "职位一:" + fmt.Sprintf("%v", po1) + ",职位二:" + fmt.Sprintf("%v", po2) + "," + fmt.Sprintf("%v", po1),
	})
	err = h.AgentService.GetPositionDifference(messages, h.AgentService.ChatModel, c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true})
}
func (h *AgentHandler) GenerateResumeInsight(c *gin.Context) {
	h.AgentService.ParseResume(c)
}

func (h *AgentHandler) GetResumeRadar(c *gin.Context) {
	var req dto.ResumeRadarReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	radar, err := h.AgentService.GetResumeRadar(c, req.UserID, req.Job)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{"success": true, "data": radar})
}
