package registryimpl

import (
	"context"
	"errors"
	"sync"
)

// use map manage plugins

type PluginMgr struct {
	pluginsMap map[string]Registry
	lock       sync.Mutex
}

// init PluginMgr
var (
	pluginMgr = &PluginMgr{
		pluginsMap: make(map[string]Registry),
	}
)

// register plugin
func RegisterPlugin(registry Registry) error {
	return pluginMgr.registerPlugin(registry)
}

func (p *PluginMgr) registerPlugin(registry Registry) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	// map list has, return
	_, ok := p.pluginsMap[registry.Name()]
	if ok {
		return errors.New("the plugin exist")
	}

	p.pluginsMap[registry.Name()] = registry
	return nil
}

// init plugin
func InitRegistry(ctx context.Context, name string, opts ...Option) (Registry, error) {
	return pluginMgr.initPlugin(ctx, name, opts...)
}

func (p *PluginMgr) initPlugin(ctx context.Context, name string, opts ...Option) (Registry, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	// map no has, return init err
	plugin, ok := p.pluginsMap[name]
	if !ok {
		return nil, errors.New("the plugin no exist, not init")
	}

	err := plugin.Init(ctx, opts...)
	return plugin, err
}
