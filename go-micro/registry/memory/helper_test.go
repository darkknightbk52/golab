package memory

import (
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
					ID:      "node-11",
					Address: "127.0.0.11",
				},
				{
					ID:      "node-12",
					Address: "127.0.0.12",
				},
			},
		},
		{
			Name:    "service-1",
			Version: "2.0",
			Nodes: []*registry.Node{
				{
					ID:      "node-21",
					Address: "127.0.0.21",
				},
				{
					ID:      "node-22",
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
					ID:      "node-1",
					Address: "127.0.0.111",
				},
				{
					ID:      "node-3",
					Address: "127.0.0.13",
				},
			},
		},
		{
			Name:    "service-1",
			Version: "3.0",
			Nodes: []*registry.Node{
				{
					ID:      "node-31",
					Address: "127.0.0.31",
				},
				{
					ID:      "node-32",
					Address: "127.0.0.32",
				},
			},
		},
	}

	srvs := addServices(olds, news)
	Expect(len(srvs)).Should(Equal(3))
	Expect()
}
