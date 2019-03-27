package micro

import (
	"context"
	"github.com/darkknightbk52/golab/go-micro/broker"
	"github.com/darkknightbk52/golab/go-micro/client"
	"github.com/darkknightbk52/golab/go-micro/registry"
	"github.com/darkknightbk52/golab/go-micro/selector"
	"github.com/darkknightbk52/golab/go-micro/server"
	"github.com/darkknightbk52/golab/go-micro/transport"
)

type Options struct {
	Context   context.Context
	Client    client.Client
	Server    server.Server
	Broker    broker.Broker
	Registry  registry.Registry
	Transport transport.Transport
}

func Context(ctx context.Context) Option {
	return func(opts *Options) {
		opts.Context = ctx
	}
}

func Client(c client.Client) Option {
	return func(opts *Options) {
		opts.Client = c
	}
}

func Server(s server.Server) Option {
	return func(opts *Options) {
		opts.Server = s
	}
}

func Broker(b broker.Broker) Option {
	return func(opts *Options) {
		opts.Broker = b
		opts.Client.Init(client.Broker(b))
		opts.Server.Init(server.Broker(b))
	}
}

func Registry(r registry.Registry) Option {
	return func(opts *Options) {
		opts.Registry = r
		opts.Client.Init(client.Registry(r))
		opts.Server.Init(server.Registry(r))
		opts.Broker.Init(broker.Registry(r))
	}
}

func Transport(t transport.Transport) Option {
	return func(opts *Options) {
		opts.Transport = t
		opts.Client.Init(client.Transport(t))
		opts.Server.Init(server.Transport(t))
	}
}

func Selector(s selector.Selector) Option {
	return func(opts *Options) {
		opts.Client.Options().Selector = s
	}
}
