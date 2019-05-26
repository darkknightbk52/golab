package memory

import (
	"context"
	"github.com/darkknightbk52/golab/go-micro/registry"
	"sync"
	"time"
)

const timeout = time.Second

type Registry struct {
	options registry.Options

	sync.RWMutex
	services map[string][]*registry.Service
	watchers map[string]*Watcher
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	options := registry.Options{
		Context: context.Background(),
	}

	for _, opt := range opts {
		opt(&options)
	}

	services := getServices(options.Context)
	if services == nil {
		services = make(map[string][]*registry.Service)
	}

	return &Registry{
		options:  options,
		services: services,
		watchers: make(map[string]*Watcher),
	}
}

func (m *Registry) Init(opts ...registry.Option) error {
	for _, o := range opts {
		o(&m.options)
	}

	services := getServices(m.options.Context)
	for k, v := range services {
		m.services[k] = addServices(m.services[k], v)
	}

	return nil
}

func (m *Registry) Options() registry.Options {
	return m.options
}

func (m *Registry) Register(s *registry.Service, opts ...registry.RegisterOption) error {
	go func() { m.watch(&registry.Result{Action: "Update", Service: s}) }()

	m.Lock()
	m.services[s.Name] = addServices(m.services[s.Name], []*registry.Service{s})
	m.Unlock()

	return nil
}

func (m *Registry) Deregister(s *registry.Service) error {
	go func() { m.watch(&registry.Result{Action: "Delete", Service: s}) }()

	m.Lock()
	m.services[s.Name] = delServices(m.services[s.Name], []*registry.Service{s})
	m.Unlock()

	return nil
}

func (m *Registry) GetService(service string) ([]*registry.Service, error) {
	m.RLock()
	defer m.RUnlock()
	return m.services[service], nil
}

func (m *Registry) ListService() ([]*registry.Service, error) {
	m.RLock()
	defer m.RUnlock()
	var services []*registry.Service
	for _, s := range m.services {
		services = append(services, s...)
	}
	return services, nil
}

func (m *Registry) Watch(opts ...registry.WatchOption) (registry.Watcher, error) {
	w := NewWatcher(opts...)
	m.Lock()
	m.watchers[w.id] = w
	m.Unlock()
	return w, nil
}

func (m *Registry) String() string {
	return "memory"
}

func (m *Registry) watch(r *registry.Result) {
	var watchers []*Watcher

	m.RLock()
	for _, w := range m.watchers {
		watchers = append(watchers, w)
	}
	m.RUnlock()

	for _, w := range watchers {
		select {
		case <-w.exit:
			m.Lock()
			delete(m.watchers, w.id)
			m.Unlock()
		default:
			select {
			case w.res <- r:
			case <-time.After(timeout):
			}
		}
	}
}

func (m Registry) Setup() {

}
