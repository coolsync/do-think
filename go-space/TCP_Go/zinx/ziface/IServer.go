package ziface

// 在ziface下创建服务模块抽象层iserver.go
type IServer interface {
	// start server
	Start()
	// stop server
	Stop()
	// run server
	Serve()
}