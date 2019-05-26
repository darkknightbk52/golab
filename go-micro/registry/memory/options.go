package memory

import (
	"context"
	"github.com/darkknightbk52/golab/go-micro/registry"
)

type serviceKey struct{}

func getServices(ctx context.Context) map[string][]*registry.Service {
	s, ok := ctx.Value(serviceKey{}).(map[string][]*registry.Service)
	if !ok {
		return nil
	}
	return s
}

func Services(s map[string]*registry.Service) registry.Option {
	return func(opts *registry.Options) {
		if opts.Context == nil {
			opts.Context = context.Background()
		}
		opts.Context = context.WithValue(opts.Context, serviceKey{}, s)
	}
}
