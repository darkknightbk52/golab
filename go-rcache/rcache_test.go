package rcache

import (
	"github.com/darkknightbk52/golab/go-micro/registry"
	"github.com/darkknightbk52/golab/go-micro/registry/memory"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func TestRCache(t *testing.T) {
	RegisterTestingT(t)

	// Build registry with some services
	memReg := memory.NewRegistry()
	err := memReg.Register(&registry.Service{
		Name:    "foo",
		Version: "1.0",
		Nodes: []*registry.Node{
			{
				Id: "1",
			},
		},
	})
	Expect(err).Should(Succeed())

	err = memReg.Register(&registry.Service{
		Name:    "foo",
		Version: "2.0",
		Nodes: []*registry.Node{
			{
				Id: "1",
			},
		},
	})
	Expect(err).Should(Succeed())

	// Wrap an empty cache onto this registry
	cacheReg := NewCache(memReg)

	// Try to get an existing service and check it be cached on
	services, err := cacheReg.GetService("foo")
	Expect(err).Should(Succeed())
	Expect(len(services)).Should(Equal(2))
	Expect(len(cacheReg.(*cache).services["foo"])).Should(Equal(2))

	// Register new service & check it be added on cache
	err = memReg.Register(&registry.Service{
		Name:    "bar",
		Version: "1.0",
		Nodes: []*registry.Node{
			{
				Id: "1",
			},
		},
	})
	Expect(err).Should(Succeed())
	time.Sleep(time.Millisecond)
	services, err = cacheReg.GetService("bar")
	Expect(err).Should(Succeed())
	Expect(len(services)).Should(Equal(1))
	Expect(len(cacheReg.(*cache).services["bar"])).Should(Equal(1))

	// Update an existing service & check it be updated on cache
	err = memReg.Register(&registry.Service{
		Name:    "foo",
		Version: "3.0",
		Nodes: []*registry.Node{
			{
				Id: "1",
			},
		},
	})
	Expect(err).Should(Succeed())
	time.Sleep(time.Millisecond)
	services, err = cacheReg.GetService("foo")
	Expect(err).Should(Succeed())
	Expect(len(services)).Should(Equal(3))
	Expect(len(cacheReg.(*cache).services["foo"])).Should(Equal(3))

	// Unregister an existing service & check it be deleted on cache
	err = memReg.Deregister(&registry.Service{
		Name:    "foo",
		Version: "3.0",
		Nodes: []*registry.Node{
			{
				Id: "1",
			},
		},
	})
	Expect(err).Should(Succeed())
	time.Sleep(time.Millisecond)
	services, err = cacheReg.GetService("foo")
	Expect(err).Should(Succeed())
	Expect(len(services)).Should(Equal(2))
	Expect(len(cacheReg.(*cache).services["foo"])).Should(Equal(2))

	err = memReg.Deregister(&registry.Service{
		Name:    "bar",
		Version: "2.0",
		Nodes: []*registry.Node{
			{
				Id: "1",
			},
		},
	})
	Expect(err).Should(Succeed())

	err = memReg.Deregister(&registry.Service{
		Name:    "bar",
		Version: "1.0",
		Nodes: []*registry.Node{
			{
				Id: "2",
			},
		},
	})
	Expect(err).Should(Succeed())
	time.Sleep(time.Millisecond)
	services, err = cacheReg.GetService("bar")
	Expect(err).Should(Succeed())
	Expect(len(services)).Should(Equal(1))
	Expect(len(cacheReg.(*cache).services["bar"])).Should(Equal(1))

	err = memReg.Deregister(&registry.Service{
		Name:    "bar",
		Version: "1.0",
		Nodes: []*registry.Node{
			{
				Id: "1",
			},
		},
	})
	Expect(err).Should(Succeed())
	time.Sleep(time.Millisecond)
	services, err = cacheReg.GetService("bar")
	Expect(err).Should(Equal(registry.ErrNotFound))

	cacheReg.Stop()
}
