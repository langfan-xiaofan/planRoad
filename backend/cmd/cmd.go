package cmd

import (
	"backend/internal/cloudflare"
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/router"
	"backend/pkg/agent"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() {
	fileParser := agent.OpenAI()
	agent := agent.NewAgent()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	agentConfig := eino.ChatModelConfig{
		APIKey:  os.Getenv("API_KEY"),
		Model:   os.Getenv("Model"),
		BaseURL: os.Getenv("BaseURL"),
	}
	client, err := eino.NewChatModel(context.Background(), &agentConfig)
	if err != nil {
		panic(err)
	}
	tools := agent.CreateTool()
	toolInfos := make([]*schema.ToolInfo, 0)
	for _, tool := range tools {
		info, err := tool.Info(context.Background())
		if err != nil {
			panic(err)
		}
		toolInfos = append(toolInfos, info)
	}
	client.BindTools(toolInfos)
	agent, err := react.NewAgent(context.Background(), &react.AgentConfig{
		Model: client,
		ToolsConfig: compose.ToolsNodeConfig{
			Tools: tools,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	config.Init()
	db.Neo4jInit()
	db.MysqlInit()
	db.RedisInit()
	db.MongoDbInit()
	r := gin.Default()
	cloudflare.CloudFlareInit()
	router.InitRouter(r, agent)
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
