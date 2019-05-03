package broker

import "github.com/darkknightbk52/golab/go-micro/registry"

type Options struct {
	Registry registry.Registry
}

type PublishOptions struct {
}

type SubscribeOptions struct {
}

type Option func(*Options)

type SubscribeOption func(*SubscribeOptions)

type PublishOption func(*PublishOptions)

func Registry(registry registry.Registry) Option {
	return func(opts *Options) {
		opts.Registry = registry
	}
}
