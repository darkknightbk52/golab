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

func addNodes(olds, news []*registry.Node) []*registry.Node {
	for _, n := range news {
		var seen bool
		for i, o := range olds {
			if o.Id == n.Id {
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

func delServices(olds, dels []*registry.Service) []*registry.Service {
	var services []*registry.Service
	for _, o := range olds {
		rm := false
		for _, d := range dels {
			if o.Version == d.Version {
				o.Nodes = delNodes(o.Nodes, d.Nodes)
				if len(o.Nodes) == 0 {
					rm = true
				}
				break
			}
		}
		if !rm {
			services = append(services, o)
		}
	}
	return services
}

func delNodes(olds, dels []*registry.Node) []*registry.Node {
	var nodes []*registry.Node
	for _, o := range olds {
		rm := false
		for _, d := range dels {
			if o.Id == d.Id {
				rm = true
				break
			}
		}
		if !rm {
			nodes = append(nodes, o)
		}
	}
	return nodes
}
