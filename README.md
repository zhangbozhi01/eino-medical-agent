# eino-medical-agent

基于 CloudWeGo **Eino** 构建的医疗 AI Agent（教育/参考用途，**非医疗设备，不能替代专业医生**）。

## 特性

- **双模式支持**
  - **自主 Agent 模式**：基于 ReAct 的 ChatModelAgent，自动收集信息、分析意图、对比治疗方案利弊
  - **工作流模式**：基于 Graph 编排的固定流程（安全分诊 → 鉴别诊断 → 方案对比 → 综合回答）
- **医学工具**（参考 Python 医疗智能体实现）
  - `safety_triage`：紧急情况关键词检测（胸痛、突然头痛、呼吸困难等）
  - `differential_diagnosis`：鉴别诊断辅助
  - `compare_treatments`：治疗方案优缺点对比
  - `check_drug_interaction`：药物相互作用检查（简化版）
- **工具扩展**
  - 官方 BrowserUse Tool（浏览器操作，查阅最新指南）
  - 可扩展 MCP（Playwright、Desktop 等）
- **安全优先**：强制免责声明、紧急情况优先检测

## 快速开始

```bash
# 配置环境变量
export OPENAI_API_KEY=your-key
export OPENAI_MODEL=gpt-4o

# 自主 Agent 模式
go run ./cmd/agent

# 工作流模式
go run ./cmd/agent -mode workflow
```

## 项目结构

```
.
├── cmd/agent/              # 入口（支持 -mode agent|workflow）
├── internal/
│   ├── agent/              # 自主 Agent（ChatModelAgent）
│   ├── workflow/           # Graph 工作流
│   └── tools/
│       ├── tools.go        # 工具加载
│       └── medical.go      # 医学专用工具
├── .env.example
└── README.md
```

## 重要声明

本项目仅供学习与研究使用。所有输出均带有“仅供参考，请咨询专业医生”的提示。检测到紧急情况时会建议立即就医。

## 参考

- [Eino](https://github.com/cloudwego/eino)
- [eino-ext](https://github.com/cloudwego/eino-ext)
- Python 参考项目：nimra-pixel/medagent、aathifpm/medAgent 等 LangGraph 医疗多智能体
