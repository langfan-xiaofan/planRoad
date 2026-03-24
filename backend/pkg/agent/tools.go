package agent

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

func GetTime(_ context.Context) string {
	return fmt.Sprintf("%s", time.Now())
}

func CreateTool() []tool.BaseTool {
	GetTimeTool := utils.NewTool(
		&schema.ToolInfo{
			Name: "get_time",
			Desc: "获取当前时间",
		}, func(ctx context.Context, _struct struct{}) (string, error) {
			return GetTime(ctx), nil
		},
	)
	return []tool.BaseTool{
		GetTimeTool,
	}
}
