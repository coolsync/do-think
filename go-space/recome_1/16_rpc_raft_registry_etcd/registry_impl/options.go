package registryimpl

import "time"

// params options design mod
type Options struct {
	Addrs        []string
	Timeout      time.Duration
	HeartBeat    int64  // 心跳check
	RegistryPath string // /path/a/b/100.xxx
}

// def func val
type Option func(opts *Options)

// options method
func WithAddrs(addrs []string) Option {
	return func(opts *Options) {
		opts.Addrs = addrs
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

func WithHeartBeat(heartBeat int64) Option {
	return func(opts *Options) {
		opts.HeartBeat = heartBeat
	}
}

func WithRegistryPath(registryPath string) Option {
	return func(opts *Options) {
		opts.RegistryPath = registryPath
	}
}
