package client

import (
	"context"
	"github.com/darkknightbk52/golab/go-micro/broker"
	"github.com/darkknightbk52/golab/go-micro/registry"
	"github.com/darkknightbk52/golab/go-micro/selector"
	"github.com/darkknightbk52/golab/go-micro/transport"
)

type Options struct {
	Broker    broker.Broker
	Registry  registry.Registry
	Transport transport.Transport
	Selector  selector.Selector
}

type Option func(*Options)

type Request interface {
}

type ReqOptions struct {
}

type RequestOption func(*ReqOptions)

type Response interface {
}

type CallOptions struct {
}

type CallOption func(*CallOptions)

type Message interface {
}

type MessageOptions struct {
}

type MessageOption func(*MessageOptions)

type Stream interface {
}

type PublishOptions struct {
}

type PublishOption func(*PublishOptions)

type Client interface {
	Init(...Option) error
	Options() Options
	NewRequest(service string, endpoint string, req interface{}, opts ...RequestOption) Request
	Call(context context.Context, req Request, rsp interface{}, opts ...CallOption) error
	Stream(context context.Context, req Request, opts ...CallOption) (Stream, error)
	NewMessage(service string, topic string, req interface{}, opts ...MessageOption) Message
	Publish(context context.Context, msg Message, opts ...PublishOption) error
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

func Selector(selector selector.Selector) Option {
	return func(opts *Options) {
		opts.Selector = selector
	}
}
