package etcd

import (
	"context"
	registry "day16/registry_project"
	"encoding/json"
	"fmt"
	"path"

	"go.etcd.io/etcd/clientv3"
)

func (e *EtcdRegistry) GetService(ctx context.Context,
	name string) (service *registry.Service, err error) {
	service, ok := e.getServiceFromCache(ctx, name)
	if ok {
		return
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	service, ok = e.getServiceFromCache(ctx, name)
	if ok {
		return
	}
	key := e.servicePath(name)
	resp, err := e.client.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return
	}
	service = &registry.Service{
		Name: name,
	}
	for _, kv := range resp.Kvs {
		value := kv.Value
		var tmpService registry.Service
		err = json.Unmarshal(value, &tmpService)
		if err != nil {
			return
		}
		for _, node := range tmpService.Nodes {
			service.Nodes = append(service.Nodes, node)
		}
	}
	allServiceInfoOld := e.value.Load().(*AllServiceInfo)
	var allServiceInfoNew = &AllServiceInfo{
		serviceMap: make(map[string]*registry.Service, MaxServiceNum),
	}
	for key, val := range allServiceInfoOld.serviceMap {
		allServiceInfoNew.serviceMap[key] = val
	}
	allServiceInfoNew.serviceMap[name] = service
	e.value.Store(allServiceInfoNew)
	return
}

func (e *EtcdRegistry) servicePath(name string) string {
	return path.Join(e.options.RegistryPath, name)
}

func (e *EtcdRegistry) registerOrKeepAlive() {
	for _, registryService := range e.registryServiceMap {
		if registryService.registered {
			e.keepAlive(registryService)
			continue
		}
		err := e.registerService(registryService)
		if err != nil {
			fmt.Println("register err")
		}
	}
}

func (e *EtcdRegistry) keepAlive(registryService *RegisterService) {
	select {
	case resp := <-registryService.keepAliveCh:
		if resp == nil {
			registryService.registered = false
			return
		}
	}
	return
}
