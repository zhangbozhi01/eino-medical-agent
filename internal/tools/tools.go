package tools

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino-ext/components/tool/browseruse"
)

// LoadAll 加载所有工具：浏览器、MCP、医学专用工具
func LoadAll(ctx context.Context) ([]tool.BaseTool, error) {
	var tools []tool.BaseTool

	// 1. 官方 BrowserUse Tool
	browserTool, err := browseruse.NewBrowserUseTool(ctx, &browseruse.Config{})
	if err != nil {
		fmt.Printf("警告: 浏览器工具初始化失败: %v\n", err)
	} else {
		tools = append(tools, browserTool)
	}

	// 2. MCP 工具（示例：可接入 Playwright / Desktop / Shell MCP）
	// mcpTools, err := loadMCPTools(ctx)
	// if err == nil {
	// 	tools = append(tools, mcpTools...)
	// }

	// 3. 医学专用工具（预留）
	// tools = append(tools, NewMedicalKnowledgeTool(), NewReportAnalysisTool())

	if len(tools) == 0 {
		return nil, fmt.Errorf("没有可用的工具")
	}
	return tools, nil
}
