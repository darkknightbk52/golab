package micro

import (
	"context"
	"github.com/darkknightbk52/golab/go-micro/client"
	"github.com/darkknightbk52/golab/go-micro/server"
)

type Option func(*Options)

type Service interface {
	Init(...Option) error
	Options() Options
	Client() client.Client
	Server() server.Server
	Run() error
	String() string
}

func NewService(opts ...Option) Service {
	return nil
}

func FromContext(ctx context.Context) (Service, bool) {
	return nil, false
}

type Function interface {
	Service
	Done() error
	Handle(v interface{}) error
	Subscribe(topic string, v interface{}) error
}

func NewFunction(opts ...Option) Function {
	return nil
}

type Publisher interface {
	Publish(ctx context.Context, msg interface{}, opts ...client.PublishOption) error
}

func NewPublisher(topic string, c client.Client) Publisher {
	return nil
}

func RegisterHandler(s server.Server, h interface{}, opts ...server.HandlerOption) error {
	return nil
}

func RegisterSubscribe(topic string, s server.Server, h interface{}, opts ...server.SubscriberOption) error {
	return nil
}
