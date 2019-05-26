package registry

import "context"

type Options struct {
	Context context.Context
}

type WatchOptions struct {
	Service string
	Context context.Context
}

func WatchService(service string) WatchOption {
	return func(opts *WatchOptions) {
		opts.Service = service
	}
}

type RegisterOptions struct {
}
