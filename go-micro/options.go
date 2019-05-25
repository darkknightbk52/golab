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
	Context     context.Context
	Name        string
	Client      client.Client
	Server      server.Server
	Broker      broker.Broker
	Registry    registry.Registry
	Transport   transport.Transport
	BeforeStart []func() error
	AfterStart  []func() error
	BeforeStop  []func() error
	AfterStop   []func() error
}

func Context(ctx context.Context) Option {
	return func(opts *Options) {
		opts.Context = ctx
	}
}

func Name(name string) Option {
	return func(opts *Options) {
		opts.Name = name
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
		//opts.Client.Options().Selector = nil
	}
}

func BeforeStart(fn func() error) Option {
	return func(opts *Options) {
		opts.BeforeStart = append(opts.BeforeStart, fn)
	}
}

func AfterStart(fn func() error) Option {
	return func(opts *Options) {
		opts.AfterStart = append(opts.AfterStart, fn)
	}
}

func BeforeStop(fn func() error) Option {
	return func(opts *Options) {
		opts.BeforeStop = append(opts.BeforeStop, fn)
	}
}

func AfterStop(fn func() error) Option {
	return func(opts *Options) {
		opts.AfterStop = append(opts.AfterStop, fn)
	}
}
