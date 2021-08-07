package main

import (
	"github.com/PGshen/go-executor-example/app"
	"github.com/PGshen/go-xxl-executor/executor"
	"github.com/PGshen/go-xxl-executor/handler"
)

func main() {
	// 初始化配置，这里根据自己的应用，可以用配置文件加载
	xxlJobConfig := executor.XxlJobConfig{
		Env:              "dev",
		AdminAddress:     "http://127.0.0.1:8080/xxl-job-admin",
		AccessToken:      "",
		Appname:          "go-executor-sample",
		Address:          "",
		Ip:               "",
		Port:             9997,
		LogPath:          "/Users/shen/Me/Study/Operation/Go/go-executor-example/log",
		LogRetentionDays: 7,
		HttpTimeout:      5,
	}
	// 注册JobHandler
	_ = handler.AddJobHandler("test", &app.ExampleJobHandler{})
	// 在实例化executor之前不可使用common.Log打印日志，因为还未读取到配置，也就还没初始化logger
	xxlExecutor := executor.NewXxlJobExecutor(xxlJobConfig)
	xxlExecutor.Start() // 启动执行器服务
}
