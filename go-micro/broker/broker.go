package broker

import "github.com/darkknightbk52/golab/go-micro/registry"

type Options struct {
	Registry registry.Registry
}

type Option func(*Options)

type Message struct {
	Header map[string]string
	Body   []byte
}

type PublishOptions struct {
}

type PublishOption func(*PublishOptions)

type Publication interface {
}

type Handler func(*Publication)

type SubscribeOptions struct {
}

type SubscribeOption func(*SubscribeOptions)

type Subscriber interface {
	Topic() string
	Message() *Message
	Ack() error
}

type Broker interface {
	Init(...Option) error
	Options() Options
	Address() string
	String() string
	Connect() error
	Disconnect() error
	Publish(string, Message, ...PublishOption) error
	Subscribe(string, Handler, ...SubscribeOption) (Subscriber, error)
}

func Registry(registry registry.Registry) Option {
	return func(opts *Options) {
		opts.Registry = registry
	}
}
