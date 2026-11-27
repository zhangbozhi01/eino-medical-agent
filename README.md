# eino-medical-agent

基于 CloudWeGo **Eino** 构建的医疗 AI Agent。

## 特性

- **双模式支持**
  - **自主 Agent 模式**：基于 ReAct 的 ChatModelAgent，适合开放式问答、研究、报告解读
  - **工作流模式**：基于 Graph 编排，适合有明确步骤的诊疗流程（症状收集 → 鉴别诊断 → 检查建议 → 方案对比）
- **工具扩展**
  - 官方 BrowserUse Tool（浏览器操作）
  - MCP 支持（可接入 Playwright MCP、Desktop MCP、Shell MCP 等）
- **医疗场景友好**：预留医学知识检索、报告解读、治疗方案对比等工具接口

## 快速开始

```bash
# 配置环境变量
export OPENAI_API_KEY=your-key
export OPENAI_MODEL=gpt-4o

# 运行自主模式
go run ./cmd/agent

# 运行工作流模式（示例）
go run ./cmd/agent -mode workflow
```

## 项目结构

```
.
├── cmd/agent/          # 入口
├── internal/
│   ├── agent/          # 自主 Agent
│   ├── workflow/       # 工作流编排
│   └── tools/          # 工具（浏览器、MCP、医学工具）
├── config/             # 配置
└── README.md
```

## 下一步

1. 完善医学专用工具（知识库检索、报告解析等）
2. 接入真实 MCP Server（浏览器 / 桌面）
3. 增加 Human-in-the-loop 审批节点
4. 对接医院内部系统

## 参考

- [Eino](https://github.com/cloudwego/eino)
- [eino-ext](https://github.com/cloudwego/eino-ext)
- [BrowserUse Tool](https://github.com/cloudwego/eino-ext/tree/main/components/tool/browseruse)
