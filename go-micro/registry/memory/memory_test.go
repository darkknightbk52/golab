package memory

import (
	"fmt"
	"github.com/darkknightbk52/golab/go-micro/registry"
	. "github.com/onsi/gomega"
	"testing"
)

var (
	testData = map[string][]*registry.Service{
		"foo": {
			{
				Name:    "foo",
				Version: "1.0.0",
				Nodes: []*registry.Node{
					{
						Id:      "foo-1.0.0-123",
						Address: "localhost",
						Port:    9999,
					},
					{
						Id:      "foo-1.0.0-321",
						Address: "localhost",
						Port:    9999,
					},
				},
			},
			{
				Name:    "foo",
				Version: "1.0.1",
				Nodes: []*registry.Node{
					{
						Id:      "foo-1.0.1-321",
						Address: "localhost",
						Port:    6666,
					},
				},
			},
			{
				Name:    "foo",
				Version: "1.0.3",
				Nodes: []*registry.Node{
					{
						Id:      "foo-1.0.3-345",
						Address: "localhost",
						Port:    8888,
					},
				},
			},
		},
		"bar": {
			{
				Name:    "bar",
				Version: "default",
				Nodes: []*registry.Node{
					{
						Id:      "bar-1.0.0-123",
						Address: "localhost",
						Port:    9999,
					},
					{
						Id:      "bar-1.0.0-321",
						Address: "localhost",
						Port:    9999,
					},
				},
			},
			{
				Name:    "bar",
				Version: "latest",
				Nodes: []*registry.Node{
					{
						Id:      "bar-1.0.1-321",
						Address: "localhost",
						Port:    6666,
					},
				},
			},
		},
	}
)

func TestMemoryRegistry(t *testing.T) {
	RegisterTestingT(t)
	m := NewRegistry()

	fn := func(k string, v []*registry.Service) {
		services, err := m.GetService(k)
		Expect(err).Should(Succeed())
		Expect(len(services)).Should(Equal(len(v)), fmt.Sprintf("Expect %d services of %s, got %d", len(services), k, len(v)))
		for _, e := range v {
			seen := false
			for _, s := range services {
				if s.Version == e.Version {
					seen = true
					break
				}
			}
			Expect(seen).Should(BeTrue(), fmt.Sprintf("Expected service %v not found", e))
		}
	}

	for _, v := range testData {
		for _, s := range v {
			err := m.Register(s)
			Expect(err).Should(Succeed())
			t.Log("Registered service:", s)
		}
	}

	for k, v := range testData {
		fn(k, v)
	}

	for _, v := range testData {
		for _, s := range v {
			err := m.Deregister(s)
			Expect(err).Should(Succeed())
			t.Log("Deregistered service:", s)
		}
	}
}
