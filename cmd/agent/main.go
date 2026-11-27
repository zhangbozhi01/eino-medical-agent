package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/zhangbozhi01/eino-medical-agent/internal/agent"
	"github.com/zhangbozhi01/eino-medical-agent/internal/workflow"
)

func main() {
	mode := flag.String("mode", "agent", "运行模式: agent | workflow")
	flag.Parse()

	ctx := context.Background()

	switch *mode {
	case "agent":
		fmt.Println("启动自主 Agent 模式...")
		if err := agent.Run(ctx); err != nil {
			log.Fatal(err)
		}
	case "workflow":
		fmt.Println("启动工作流模式...")
		if err := workflow.Run(ctx); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("未知模式，请使用 -mode agent 或 -mode workflow")
		os.Exit(1)
	}
}
