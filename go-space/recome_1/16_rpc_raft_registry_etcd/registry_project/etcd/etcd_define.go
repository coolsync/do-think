package etcd

import (
	"sync"
	"sync/atomic"

	registry "day16/registry_project"

	"go.etcd.io/etcd/clientv3"
)

type AllServiceInfo struct {
	serviceMap map[string]*registry.Service
}

type RegisterService struct {
	id          clientv3.LeaseID
	service     *registry.Service
	registered  bool
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
}

type EtcdRegistry struct {
	options            *registry.Options
	client             *clientv3.Client
	serviceCh          chan *registry.Service
	value              atomic.Value
	lock               sync.Mutex
	registryServiceMap map[string]*RegisterService
}

const (
	MaxServiceNum = 8
)

var (
	etcdRegistry *EtcdRegistry = &EtcdRegistry{
		serviceCh:          make(chan *registry.Service, MaxServiceNum),
		registryServiceMap: make(map[string]*RegisterService, MaxServiceNum),
	}
)
