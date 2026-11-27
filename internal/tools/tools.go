package tools

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino-ext/components/tool/browseruse"
)

// LoadAll 加载所有工具：浏览器、医学专用工具
func LoadAll(ctx context.Context) ([]tool.BaseTool, error) {
	var tools []tool.BaseTool

	// 医学工具（优先）
	tools = append(tools, NewMedicalTools()...)

	// 浏览器工具（可选，失败不阻断）
	browserTool, err := browseruse.NewBrowserUseTool(ctx, &browseruse.Config{})
	if err != nil {
		fmt.Printf("警告: 浏览器工具初始化失败（可能缺少 Chrome）: %v\n", err)
	} else {
		tools = append(tools, browserTool)
	}

	if len(tools) == 0 {
		return nil, fmt.Errorf("没有可用的工具")
	}
	return tools, nil
}
