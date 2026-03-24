package service

import (
	"backend/internal/dto"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/gin-gonic/gin"
)

type AgentService struct {
	Client *react.Agent
}

func NewAgentService(client *react.Agent) *AgentService {
	return &AgentService{
		Client: client,
	}
}

func (s *AgentService) Chat(req *dto.ChatRequest, message []*schema.Message, c *gin.Context) error {
	resp, err := s.Client.Stream(context.Background(), message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	messchan := make(chan string)
	go Stream(c, messchan)
	for {
		mess, err := resp.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		messchan <- mess.Content
		fmt.Println(mess.Content)
	}
	return nil
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
