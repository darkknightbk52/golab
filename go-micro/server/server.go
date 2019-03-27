package server

import (
	"github.com/darkknightbk52/golab/go-micro/broker"
	"github.com/darkknightbk52/golab/go-micro/registry"
	"github.com/darkknightbk52/golab/go-micro/transport"
)

type Handler interface {
}

type HandlerOptions struct {
}

type HandlerOption func(*HandlerOptions)

type Options struct {
	Broker    broker.Broker
	Registry  registry.Registry
	Transport transport.Transport
}

type Option func(*Options)

type Subscriber interface {
}

type SubscriberOptions struct {
}

type SubscriberOption func(*SubscriberOptions)

type Server interface {
	Init(...Option) error
	Options() Options
	Start() error
	Stop() error
	NewHandler(interface{}, ...HandlerOption) (Handler, error)
	Handle(Handler) error
	NewSubscriber(string, interface{}, ...SubscriberOption) (Subscriber, error)
	Subscribe(Subscriber) error
}

func Broker(broker broker.Broker) Option {
	return func(opts *Options) {
		opts.Broker = broker
	}
}

func Registry(registry registry.Registry) Option {
	return func(opts *Options) {
		opts.Registry = registry
	}
}

func Transport(transport transport.Transport) Option {
	return func(opts *Options) {
		opts.Transport = transport
	}
}
