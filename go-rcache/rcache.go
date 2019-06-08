package rcache

import (
	"github.com/darkknightbk52/golab/go-micro/registry"
	"sync"
	"time"
)

type Cache interface {
	registry.Registry
	Stop()
}

type cache struct {
	registry.Registry
	opts Options

	sync.RWMutex
	services map[string][]*registry.Service
	ttls     map[string]time.Time
	watched  map[string]bool

	exit chan bool
}

const DefaultTTL = time.Minute

func NewCache(r registry.Registry, opts ...Option) Cache {
	options := Options{
		TTL: DefaultTTL,
	}

	for _, o := range opts {
		o(&options)
	}

	return &cache{
		Registry: r,
		opts:     options,
		services: make(map[string][]*registry.Service),
		ttls:     make(map[string]time.Time),
		watched:  make(map[string]bool),
		exit:     make(chan bool),
	}
}

func (c *cache) Stop() {
	select {
	case <-c.exit:
		return
	default:
		close(c.exit)
	}
}

func (c *cache) GetService(service string) ([]*registry.Service, error) {
	services, err := c.get(service)
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, registry.ErrNotFound
	}

	return services, nil
}

func (c *cache) get(service string) ([]*registry.Service, error) {
	c.RLock()

	services := c.services[service]
	ttl := c.ttls[service]
	if c.isValid(services, ttl) {
		c.RUnlock()
		return c.cp(services), nil
	}

	// ask registry
	get := func(service string) ([]*registry.Service, error) {
		services, err := c.Registry.GetService(service)
		if err != nil {
			return nil, err
		}
		if len(services) == 0 {
			return nil, registry.ErrNotFound
		}

		c.Lock()
		// cache them
		c.set(service, services)
		c.Unlock()

		return services, nil
	}

	// watch recently cached service
	if _, ok := c.watched[service]; !ok {
		// monitor updates from services
		go c.run(service)
	}
	c.RUnlock()

	return get(service)
}

func (c cache) isValid(services []*registry.Service, ttl time.Time) bool {
	if len(services) == 0 {
		return false
	}

	if time.Since(ttl) > c.opts.TTL {
		return false
	}

	return true
}

func (c *cache) cp(current []*registry.Service) []*registry.Service {
	var copy []*registry.Service
	for _, svc := range current {
		s := *svc
		for _, node := range svc.Nodes {
			n := *node
			s.Nodes = append(s.Nodes, &n)
		}

		for _, endpoint := range svc.Endpoints {
			e := *endpoint
			s.Endpoints = append(s.Endpoints, &e)
		}
		copy = append(copy, &s)
	}
	return copy
}

func (c *cache) set(name string, services []*registry.Service) {
	c.services[name] = services
	c.ttls[name] = time.Now().Add(c.opts.TTL)
}

func (c *cache) run(service string) {
	c.Lock()
	c.watched[service] = true
	c.Unlock()

	defer func() {
		c.Lock()
		delete(c.watched, service)
		c.Unlock()
	}()

	for {
		if c.isQuit() {
			return
		}

		watcher, err := c.Registry.Watch(registry.WatchService(service))
		if err != nil {
			if c.isQuit() {
				return
			}

			continue
		}

		err = c.watch(watcher)
		if err != nil {
			if c.isQuit() {
				return
			}

			continue
		}
	}
}

func (c *cache) watch(watcher registry.Watcher) error {
	go func() {
		<-c.exit
		watcher.Stop()
	}()

	for {
		r, err := watcher.Next()
		if err != nil {
			return err
		}
		c.update(r)
	}
}

func (c *cache) isQuit() bool {
	select {
	case <-c.exit:
		return true
	default:
		return false
	}
}

func (c *cache) update(r *registry.Result) {
	if r == nil || len(r.Service.Name) == 0 {
		return
	}

	c.RLock()
	services, ok := c.services[r.Service.Name]
	c.RUnlock()

	if !ok {
		return
	}

	if len(r.Service.Nodes) == 0 {
		switch r.Action {
		case "delete":
			c.del(r.Service.Name)
			return
		}
	}

	var service *registry.Service
	var index int
	for i, s := range services {
		if r.Service.Version == s.Version {
			service = s
			index = i
		}
	}

	c.Lock()
	defer c.Unlock()

	switch r.Action {
	case "Create", "Update":
		if service == nil {
			c.set(r.Service.Name, append(services, r.Service))
			return
		}

		for _, cur := range service.Nodes {
			seen := false
			for _, n := range r.Service.Nodes {
				if cur.Id == n.Id {
					seen = true
					break
				}
			}
			if !seen {
				r.Service.Nodes = append(r.Service.Nodes, cur)
			}
		}

		services[index] = r.Service
		c.set(r.Service.Name, services)
	case "Delete":
		if service == nil {
			return
		}

		var nodes []*registry.Node
		for _, cur := range service.Nodes {
			seen := false
			for _, del := range r.Service.Nodes {
				if cur.Id == del.Id {
					seen = true
					break
				}
			}
			if !seen {
				nodes = append(nodes, cur)
			}
		}

		if len(nodes) > 0 {
			service.Nodes = nodes
			services[index] = service
			c.set(r.Service.Name, services)
			return
		}

		if len(services) == 1 {
			c.del(r.Service.Name)
			return
		}

		var srvs []*registry.Service
		for _, s := range services {
			if s.Version != service.Version {
				srvs = append(srvs, s)
			}
		}

		c.set(r.Service.Name, srvs)
	}
}

func (c *cache) del(service string) {
	delete(c.services, service)
	delete(c.watched, service)
}
