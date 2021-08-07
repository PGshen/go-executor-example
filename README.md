# go-executor-example

### 步骤
1. mod引入go-xxl-executor包
```shell
require github.com/PGshen/go-xxl-executor v1.0.1
```

2. 开发JobHandler业务逻辑
```go
type ExampleJobHandler struct {
	handler.MethodJobHandler
}

func (receiver *ExampleJobHandler) Execute(param handler.Param) biz.ReturnT {
	receiver.MethodJobHandler.Execute(param)
	common.Log.Info("Test...")
	return biz.NewReturnT(common.SuccessCode, "Test JobHandler")
}
```

3.  初始化配置
```go
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
```

4. 注册JobHandler到执行器并启动
```go
// 注册JobHandler
_ = handler.AddJobHandler("test", &app.ExampleJobHandler{})
// 在实例化executor之前不可使用common.Log打印日志，因为还未读取到配置，也就还没初始化logger
xxlExecutor := executor.NewXxlJobExecutor(xxlJobConfig)
xxlExecutor.Start() // 启动执行器服务
```

5. 在xxl-job-admin上添加执行器，然后添加任务