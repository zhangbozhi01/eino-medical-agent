package agent

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino-ext/components/model/openai"

	"github.com/zhangbozhi01/eino-medical-agent/internal/tools"
)

func Run(ctx context.Context) error {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("请设置 OPENAI_API_KEY 环境变量")
	}
	modelName := os.Getenv("OPENAI_MODEL")
	if modelName == "" {
		modelName = "gpt-4o"
	}

	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey: apiKey,
		Model:  modelName,
	})
	if err != nil {
		return err
	}

	allTools, err := tools.LoadAll(ctx)
	if err != nil {
		return err
	}

	ag, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        "MedicalAgent",
		Description: "专业医疗辅助智能体，支持自主推理与工具调用",
		Instruction: `你是一个专业的医疗 AI 助手。
请严格遵循循证医学原则，给出有依据的建议。
涉及诊断或治疗方案时，请明确说明“仅供参考，请咨询专业医生”。
你可以使用浏览器工具查阅最新指南，也可以调用其他扩展工具。`,
		Model: chatModel,
		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: allTools,
			},
		},
		MaxIterations: 15,
	})
	if err != nil {
		return err
	}

	runner := adk.NewRunner(ctx, adk.RunnerConfig{Agent: ag})

	query := "患者出现持续性头痛伴恶心，请帮我分析可能的原因并建议下一步检查。"
	fmt.Printf("用户: %s\n\n", query)

	iter := runner.Query(ctx, query)
	for {
		event, ok := iter.Next()
		if !ok {
			break
		}
		if event.Message != nil && event.Message.Content != "" {
			fmt.Printf("助手: %s\n", event.Message.Content)
		}
	}
	return nil
}
