package workflow

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/compose"
)

// ClinicalState 工作流状态（参考 LangGraph State）
type ClinicalState struct {
	Symptoms      string
	IsEmergency   bool
	Differentials string
	Treatment     string
	FinalAnswer   string
}

// Run 使用 Graph 编排固定诊疗流程
func Run(ctx context.Context) error {
	fmt.Println("=== 工作流模式（Graph 编排） ===")

	g := compose.NewGraph[*ClinicalState, *ClinicalState]()

	// 节点1: 安全分诊
	_ = g.AddLambdaNode("safety", compose.InvokableLambda(func(ctx context.Context, state *ClinicalState) (*ClinicalState, error) {
		fmt.Println("→ 安全分诊节点")
		// 实际应调用 SafetyTriageTool
		state.IsEmergency = false
		if state.Symptoms == "" {
			state.Symptoms = "示例：持续性头痛伴恶心"
		}
		return state, nil
	}))

	// 节点2: 鉴别诊断
	_ = g.AddLambdaNode("diagnosis", compose.InvokableLambda(func(ctx context.Context, state *ClinicalState) (*ClinicalState, error) {
		fmt.Println("→ 鉴别诊断节点")
		state.Differentials = "偏头痛、紧张性头痛、需排除颅内病变（仅供参考）"
		return state, nil
	}))

	// 节点3: 方案对比
	_ = g.AddLambdaNode("treatment", compose.InvokableLambda(func(ctx context.Context, state *ClinicalState) (*ClinicalState, error) {
		fmt.Println("→ 治疗方案对比节点")
		state.Treatment = "急性期止痛 vs 预防性治疗，各有利弊。请咨询医生。"
		return state, nil
	}))

	// 节点4: 综合输出
	_ = g.AddLambdaNode("synthesize", compose.InvokableLambda(func(ctx context.Context, state *ClinicalState) (*ClinicalState, error) {
		fmt.Println("→ 综合回答节点")
		state.FinalAnswer = fmt.Sprintf(`【仅供参考，请咨询专业医生】

症状：%s
紧急情况：%v
可能方向：%s
方案提示：%s

建议：进一步检查（如需要影像）并到神经内科就诊。`, state.Symptoms, state.IsEmergency, state.Differentials, state.Treatment)
		return state, nil
	}))

	_ = g.AddEdge(compose.START, "safety")
	_ = g.AddEdge("safety", "diagnosis")
	_ = g.AddEdge("diagnosis", "treatment")
	_ = g.AddEdge("treatment", "synthesize")
	_ = g.AddEdge("synthesize", compose.END)

	runnable, err := g.Compile(ctx)
	if err != nil {
		return err
	}

	initial := &ClinicalState{Symptoms: "持续性头痛伴恶心三天"}
	result, err := runnable.Invoke(ctx, initial)
	if err != nil {
		return err
	}

	fmt.Println("\n最终结果：")
	fmt.Println(result.FinalAnswer)
	return nil
}
