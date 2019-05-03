package memory

import (
	"github.com/darkknightbk52/golab/go-micro/broker"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"sync"
)

type memoryBroker struct {
	options     broker.Options
	mutex       sync.Mutex
	connected   bool
	subscribers map[string][]*memorySubscriber
}

func NewMemoryBroker(opts ...broker.Option) *memoryBroker {
	var options broker.Options
	for _, o := range opts {
		o(&options)
	}
	return &memoryBroker{
		options:     options,
		connected:   false,
		subscribers: make(map[string][]*memorySubscriber),
	}
}

func (b *memoryBroker) Init(opts ...broker.Option) error {
	for _, o := range opts {
		o(&b.options)
	}
	return nil
}

func (b *memoryBroker) Options() broker.Options {
	return b.options
}

func (b *memoryBroker) Address() string {
	return "memory"
}

func (b *memoryBroker) String() string {
	return "memory"
}

func (b *memoryBroker) Connect() error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if b.connected {
		return nil
	}
	b.connected = true
	return nil
}

func (b *memoryBroker) Disconnect() error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if !b.connected {
		return nil
	}
	b.connected = false
	return nil
}

func (b *memoryBroker) Publish(topic string, msg broker.Message, opts ...broker.PublishOption) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if !b.connected {
		return errors.New("broker is disconnected")
	}

	subs, ok := b.subscribers[topic]
	if !ok {
		return nil
	}

	pub := &memoryPublication{
		message: &msg,
	}
	for _, s := range subs {
		s.handle(pub)
	}
	return nil
}

func (b *memoryBroker) Subscribe(topic string, h broker.Handler, opts ...broker.SubscribeOption) (broker.Subscriber, error) {
	if !b.connected {
		return nil, errors.New("not connected")
	}

	var options broker.SubscribeOptions
	for _, o := range opts {
		o(&options)
	}
	sub := &memorySubscriber{
		exit:    make(chan bool),
		id:      uuid.New().String(),
		topic:   topic,
		handler: h,
		options: options,
	}
	b.subscribers[topic] = append(b.subscribers[topic], sub)

	go func() {
		<-sub.exit
		var newSubscribers []*memorySubscriber
		for _, s := range b.subscribers[sub.topic] {
			if s.id == sub.id {
				continue
			}
			newSubscribers = append(newSubscribers, s)
		}
		b.subscribers[sub.topic] = newSubscribers
	}()

	return sub, nil
}

type memorySubscriber struct {
	exit    chan bool
	id      string
	topic   string
	handler broker.Handler
	options broker.SubscribeOptions
}

func (s *memorySubscriber) Topic() string {
	return s.topic
}

func (s *memorySubscriber) Options() broker.SubscribeOptions {
	return s.options
}

func (s *memorySubscriber) Unsubscribe() error {
	s.exit <- true
	return nil
}

func (s *memorySubscriber) handle(pub broker.Publication) error {
	return s.handler(pub)
}

type memoryPublication struct {
	message *broker.Message
	options broker.PublishOptions
}

func (p *memoryPublication) Message() *broker.Message {
	return p.message
}

func (p *memoryPublication) Options() broker.PublishOptions {
	return p.options
}

func (p *memoryPublication) Ack() error {
	return nil
}
