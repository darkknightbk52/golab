package memory

import (
	"github.com/darkknightbk52/golab/go-micro/registry"
	"sync"
)

type Registry struct {
	options registry.Options

	sync.RWMutex
	Services map[string][]*registry.Service
	Watchers map[string]registry.Watcher
}

func (m *Registry) Init(opts ...registry.Option) error {
	for _, o := range opts {
		o(&m.options)
	}


}

func (m *Registry) Options() registry.Options {
	panic("implement me")
}

func (m *Registry) Register(*registry.Service, ...registry.RegisterOption) error {
	panic("implement me")
}

func (m *Registry) Deregister(*registry.Service) error {
	panic("implement me")
}

func (m *Registry) GetService(string) ([]*registry.Service, error) {
	panic("implement me")
}

func (m *Registry) ListService() ([]*registry.Service, error) {
	panic("implement me")
}

func (m *Registry) Watch(...registry.WatchOption) (registry.Watcher, error) {
	panic("implement me")
}

func (m *Registry) String() string {
	panic("implement me")
}

func (m Registry) Setup() {

}

func NewRegistry() registry.Registry {
	return &Registry{}
}
