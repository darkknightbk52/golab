package memory

import (
	"fmt"
	"github.com/darkknightbk52/golab/go-micro/registry"
	. "github.com/onsi/gomega"
	"testing"
)

func TestAddServices(t *testing.T) {
	RegisterTestingT(t)

	olds := []*registry.Service{
		{
			Name:    "service-1",
			Version: "1.0",
			Nodes: []*registry.Node{
				{
					Id:      "node-11",
					Address: "127.0.0.11",
				},
				{
					Id:      "node-12",
					Address: "127.0.0.12",
				},
			},
		},
		{
			Name:    "service-1",
			Version: "2.0",
			Nodes: []*registry.Node{
				{
					Id:      "node-21",
					Address: "127.0.0.21",
				},
				{
					Id:      "node-22",
					Address: "127.0.0.22",
				},
			},
		},
	}

	news := []*registry.Service{
		{
			Name:    "service-1",
			Version: "1.0",
			Nodes: []*registry.Node{
				{
					Id:      "node-11",
					Address: "127.0.0.111",
				},
				{
					Id:      "node-13",
					Address: "127.0.0.13",
				},
			},
		},
		{
			Name:    "service-1",
			Version: "3.0",
			Nodes: []*registry.Node{
				{
					Id:      "node-31",
					Address: "127.0.0.31",
				},
				{
					Id:      "node-32",
					Address: "127.0.0.32",
				},
			},
		},
	}

	services := addServices(olds, news)
	Expect(len(services)).Should(Equal(3))

	search := func(services []*registry.Service, version string) *registry.Service {
		for _, s := range services {
			if s.Version == version {
				return s
			}
		}
		return nil
	}

	expectedVersions := []string{
		"1.0",
		"2.0",
		"3.0",
	}

	for _, ev := range expectedVersions {
		s := search(services, ev)
		Expect(s).ShouldNot(BeNil(), "Not found expected version: "+ev)
		if ev == "1.0" {
			Expect(len(s.Nodes)).Should(Equal(3), fmt.Sprintf("Unexpected number of nodes: %d", len(s.Nodes)))
		} else {
			Expect(len(s.Nodes)).Should(Equal(2), fmt.Sprintf("Unexpected number of nodes: %d", len(s.Nodes)))
		}
		t.Log(s)
	}
}
