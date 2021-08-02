package registry

import (
	"context"
	"fmt"
	"sync"
)

// 插件管理类
// ▪ 可以用一个大map管理，key字符串，value是Registry接口对象
// ▪ 用户自定义去调用，自定义插件
// ▪ 实现注册中心的初始化，供系统使用

type PluginMgr struct {
	pluginMap map[string]Registry
	lock      sync.Mutex
}

// map init
var (
	pluginMgr = &PluginMgr{
		pluginMap: make(map[string]Registry),
	}
)

// plugins resigster
func RegisterPlugin(plugin Registry) (err error) {
	return pluginMgr.registerPlugin(plugin)
}

func (p *PluginMgr) registerPlugin(plugin Registry) (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	// in has if not exists
	_, ok := p.pluginMap[plugin.Name()]
	if ok {
		err = fmt.Errorf("the plugin is exist")
		return
	}
	p.pluginMap[plugin.Name()] = plugin
	return
}

// registry center init
func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return pluginMgr.initPlugin(ctx, name, opts...)
}

func (p *PluginMgr) initPlugin(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	plugin, ok := p.pluginMap[name]
	if !ok {
		err = fmt.Errorf("the plugin %s is not exist", name)
		return
	}
	registry = plugin
	err = plugin.Init(ctx, opts...)
	return
}
