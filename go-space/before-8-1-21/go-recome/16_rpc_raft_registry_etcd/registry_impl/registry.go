package registryimpl

import "context"

// 定义服务注册总接口 Registry，定义方法
// ▪ Name()：插件名，例如传etcd
// ▪ Init(opts ...Option)：初始化，里面用选项设计模式做初始化
// ▪ Register()：服务注册
// ▪ Unregister()：服务反注册，例如服务端停了，注册列表销毁
// ▪ GetService：服务发现（ip port[] string）

type Registry interface {
	Name() string                                                  // plugin name
	Init(ctx context.Context, opts ...Option) error                // Init, use options design mod
	Register(ctx context.Context, service *Service) error          // service register
	Unregister(ctx context.Context, service *Service) error        // 服务反注册
	GetService(ctx context.Context, name string) (*Service, error) // service found
}
