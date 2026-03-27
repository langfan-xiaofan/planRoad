package agent

import (
	"context"
	"fmt"
	"os"

	eino "github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func Eino() *eino.ChatModel {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	agentConfig := eino.ChatModelConfig{
		APIKey:  os.Getenv("API_KEY"),
		Model:   os.Getenv("Qwen_longModel"),
		BaseURL: os.Getenv("BaseURL"),
	}
	client, err := eino.NewChatModel(context.Background(), &agentConfig)
	if err != nil {
		panic(err)
	}
	return client
}

func CreateFileParseAgent() *eino.ChatModel {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	agentConfig := eino.ChatModelConfig{
		APIKey:  os.Getenv("API_KEY"),
		Model:   os.Getenv("Qwen_plusModel"),
		BaseURL: os.Getenv("BaseURL"),
	}
	client, err := eino.NewChatModel(context.Background(), &agentConfig)
	if err != nil {
		panic(err)
	}
	return client
}

func OpenAI() *openai.Client {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	config := openai.DefaultConfig(os.Getenv("API_KEY"))
	config.BaseURL = os.Getenv("BaseURL")
	client := openai.NewClientWithConfig(config)
	return client
}

func NewAgent() *react.Agent {
	client := Eino()
	tools := CreateTool()
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
		return nil
	}
	// chain := compose.NewChain[string, []*schema.Message]()
	// chain.AppendToolsNode()
	return agent
}
