package agent

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

type ParseFileRequest struct {
	FileNames   []string `json:"file_names" description:"文件名"`
	UserMessage string   `json:"user_message" description:"用户消息"`
}

func GetTime(_ context.Context, _struct struct{}) (string, error) {
	return fmt.Sprintf("%s", time.Now()), nil
}

// func ParseFile(_ context.Context, req ParseFileRequest) (string, error) {
// 	client := OpenAI()
// 	// var parsedFiles map[string]string = make(map[string]string)
// 	var fileIDs map[string]string = make(map[string]string)
// 	for _, fileName := range req.FileNames {
// 		uploadReq := openai.FileRequest{
// 			FilePath: fmt.Sprintf("./uploads/%s", fileName),
// 			Purpose:  "file-extract",
// 			FileName: filepath.Base(fileName),
// 		}
// 		file, err := client.CreateFile(context.Background(), uploadReq)
// 		if err != nil {
// 			fmt.Println(err)
// 			return "", err
// 		}
// 		fmt.Println("上传成功")
// 		err = os.Remove(fmt.Sprintf("./uploads/%s", fileName))
// 		if err != nil {
// 			fmt.Println(err)
// 			return "", err
// 		}
// 		fileIDs[fileName] = file.ID
// 	}
// 	requestMessage := make([]openai.ChatCompletionMessage, 0)
// 	requestMessage = append(requestMessage, openai.ChatCompletionMessage{
// 		Role:    openai.ChatMessageRoleSystem,
// 		Content: FilePrompt,
// 	})
// 	for _, fileID := range fileIDs {
// 		requestMessage = append(requestMessage, openai.ChatCompletionMessage{
// 			Role:    openai.ChatMessageRoleSystem,
// 			Content: fmt.Sprintf("fileid://%s", fileID),
// 		})
// 	}
// 	requestMessage = append(requestMessage, openai.ChatCompletionMessage{
// 		Role:    openai.ChatMessageRoleUser,
// 		Content: req.UserMessage,
// 	})
// 	messageReq := openai.ChatCompletionRequest{
// 		Model:    os.Getenv("Qwen_longModel"),
// 		Messages: requestMessage,
// 	}
// 	resp, err := client.CreateChatCompletion(context.Background(), messageReq)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}
// 	// for _, choice := range resp.Choices {
// 	// 	parsedFiles[choice.Message.Content] = choice.Message.Content
// 	// }
// 	fmt.Println(resp.Choices[0].Message.Content)
// 	return resp.Choices[0].Message.Content, nil
// }

func CreateTool() []tool.BaseTool {
	GetTimeTool := utils.NewTool(
		&schema.ToolInfo{
			Name: "get_time",
			Desc: "获取当前时间",
		}, GetTime,
	)

	return []tool.BaseTool{
		GetTimeTool,
	}
}
