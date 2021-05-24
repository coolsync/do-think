package registry

import (
	"time"
)

// 选项设计模式，实现参数初始化

// Define Options struct
type Options struct {
	Addrs        []string
	Timeout      time.Duration
	HeartBeat    int64
	RegistryPath string // /path/a/b/100.xxx
}

// define func var
type Option func(op *Options)

// Options method
func WithAddrs(addrs []string) Option {
	return func(op *Options) {
		op.Addrs = addrs
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(op *Options) {
		op.Timeout = timeout
	}
}

func WithHeartBeat(heartbeat int64) Option {
	return func(op *Options) {
		op.HeartBeat = heartbeat
	}
}
func WithRegistryPath(regpath string) Option {
	return func(op *Options) {
		op.RegistryPath = regpath
	}

}
