package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

// SafetyTriageTool 紧急情况分诊工具（参考 Python medical agents 的 SafetyAgent）
type SafetyTriageTool struct{}

func (t *SafetyTriageTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "safety_triage",
		Desc: "检测用户描述中的紧急医疗情况（如胸痛、突然严重头痛、呼吸困难、意识丧失等）。必须优先调用。",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"symptoms": {
				Type: schema.String,
				Desc: "用户描述的症状或问题",
			},
		}),
	}, nil
}

func (t *SafetyTriageTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	var args struct {
		Symptoms string `json:"symptoms"`
	}
	if err := json.Unmarshal([]byte(argumentsInJSON), &args); err != nil {
		return "", err
	}

	emergencyKeywords := []string{
		"胸痛", "胸口痛", "心绞痛", "突然头痛", "剧烈头痛", "偏瘫", "口角歪斜", "言语不清",
		"呼吸困难", "喘不上气", "意识丧失", "昏迷", "抽搐", "大量出血", "自杀", "自残",
		"chest pain", "sudden headache", "paralysis", "shortness of breath", "loss of consciousness",
	}

	text := strings.ToLower(args.Symptoms)
	var found []string
	for _, kw := range emergencyKeywords {
		if strings.Contains(text, strings.ToLower(kw)) {
			found = append(found, kw)
		}
	}

	if len(found) > 0 {
		return fmt.Sprintf(`{"is_emergency": true, "keywords": %q, "advice": "检测到可能紧急情况，请立即拨打急救电话（中国 120，美国 911）或前往最近急诊。本系统仅供参考，不能替代专业医疗。"}`, found), nil
	}
	return `{"is_emergency": false, "advice": "未检测到明显紧急关键词。请继续详细描述，但仍需咨询专业医生。"}`, nil
}

// DifferentialDiagnosisTool 鉴别诊断工具
type DifferentialDiagnosisTool struct{}

func (t *DifferentialDiagnosisTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "differential_diagnosis",
		Desc: "根据症状生成可能的鉴别诊断列表（仅供参考，非最终诊断）。",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"symptoms": {Type: schema.String, Desc: "症状描述"},
			"age":      {Type: schema.String, Desc: "年龄（可选）"},
			"gender":   {Type: schema.String, Desc: "性别（可选）"},
		}),
	}, nil
}

func (t *DifferentialDiagnosisTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	var args struct {
		Symptoms string `json:"symptoms"`
		Age      string `json:"age"`
		Gender   string `json:"gender"`
	}
	_ = json.Unmarshal([]byte(argumentsInJSON), &args)

	result := map[string]any{
		"note": "此为辅助分析，非诊断。请结合专业医生意见。",
		"possible_differentials": []string{
			"需进一步收集病史和检查",
			"建议使用浏览器工具柨询最新指南",
		},
		"next_steps": "详细询问发病时间、诱因、伴随症状、既往史、用药史",
	}
	b, _ := json.Marshal(result)
	return string(b), nil
}

// CompareTreatmentsTool 治疗方案利弊对比
type CompareTreatmentsTool struct{}

func (t *CompareTreatmentsTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "compare_treatments",
		Desc: "对比不同治疗方案的优缺点、适用人群、风险。",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"condition": {Type: schema.String, Desc: "病情或诊断"},
			"options":   {Type: schema.String, Desc: "待对比的方案，用逗号分隔"},
		}),
	}, nil
}

func (t *CompareTreatmentsTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	var args struct {
		Condition string `json:"condition"`
		Options   string `json:"options"`
	}
	_ = json.Unmarshal([]byte(argumentsInJSON), &args)

	result := map[string]any{
		"disclaimer": "仅供参考，具体方案请咨询专业医生。",
		"condition":  args.Condition,
		"comparison": "请结合最新循证医学证据对比。可使用浏览器工具查询 Cochrane、UpToDate 等。",
		"advice":     "列出每种方案的 pros/cons、证据等级、禁忌。",
	}
	b, _ := json.Marshal(result)
	return string(b), nil
}

// DrugInteractionTool 药物相互作用检查（简化版）
type DrugInteractionTool struct{}

func (t *DrugInteractionTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "check_drug_interaction",
		Desc: "检查药物之间的潜在相互作用（教育用途，非完整数据库）。",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"drugs": {Type: schema.String, Desc: "药物名称，逗号分隔"},
		}),
	}, nil
}

func (t *DrugInteractionTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	var args struct {
		Drugs string `json:"drugs"`
	}
	_ = json.Unmarshal([]byte(argumentsInJSON), &args)

	result := map[string]any{
		"disclaimer": "此为简化检查，请使用专业药物相互作用数据库（如 OpenFDA、Micromedex）。",
		"drugs":      args.Drugs,
		"note":       "建议咨询药师或医生。",
	}
	b, _ := json.Marshal(result)
	return string(b), nil
}

// NewMedicalTools 返回所有医学工具
func NewMedicalTools() []tool.BaseTool {
	return []tool.BaseTool{
		&SafetyTriageTool{},
		&DifferentialDiagnosisTool{},
		&CompareTreatmentsTool{},
		&DrugInteractionTool{},
	}
}
