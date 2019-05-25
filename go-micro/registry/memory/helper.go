package memory

import "github.com/darkknightbk52/golab/go-micro/registry"

func addServices(olds, news []*registry.Service) []*registry.Service {
	for _, n := range news {
		var seen bool
		for i, o := range olds {
			if o.Version == n.Version {
				seen = true
				n.Nodes = addNodes(o.Nodes, n.Nodes)
				olds[i] = n
			}
		}
		if !seen {
			olds = append(olds, n)
		}
	}
	return olds
}

func addNodes(olds []*registry.Node, news []*registry.Node) []*registry.Node {
	for _, n := range news {
		var seen bool
		for i, o := range olds {
			if o.ID == n.ID {
				seen = true
				olds[i] = n
			}
		}
		if !seen {
			olds = append(olds, n)
		}
	}
	return olds
}
