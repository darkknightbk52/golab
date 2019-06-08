package rcache

import "time"

type Options struct {
	TTL time.Duration
}

type Option func(*Options)

func TTLCache(ttl time.Duration) Option {
	return func(opts *Options) {
		opts.TTL = ttl
	}
}
