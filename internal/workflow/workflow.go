package workflow

import (
	"context"
	"fmt"
)

// Run 工作流模式示例
// 后续可使用 Eino Graph 实现：
// 症状收集 → 鉴别诊断 → 检查建议 → 方案对比 → 医生审核
func Run(ctx context.Context) error {
	fmt.Println("工作流模式骨架已创建。")
	fmt.Println("可在此使用 compose.NewGraph 编排固定诊疗流程。")
	fmt.Println("示例节点：")
	fmt.Println("  1. 症状收集节点")
	fmt.Println("  2. 鉴别诊断节点（可调用知识库）")
	fmt.Println("  3. 检查建议节点")
	fmt.Println("  4. 治疗方案对比节点")
	fmt.Println("  5. Human-in-the-loop 审核节点")
	return nil
}
