package selector

import "github.com/darkknightbk52/golab/go-micro/registry"

type Strategy func([]*registry.Node) Next

type Options struct {
	Registry registry.Registry
	Strategy Strategy
}

type Option func(*Options)

type SelectOptions struct {
}

type SelectOption func(SelectOptions)

type Next func() (*registry.Node, error)

type Selector interface {
	Init(...Option) error
	Options() Options
	Close() error
	Select(service string, opts ...SelectOption) (Next, error)
	Mark(service string, node *registry.Node, err error)
	Reset(service string)
}
