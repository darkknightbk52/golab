package broker

type Handler func(Publication) error

type Publication interface {
	Message() *Message
	Options() PublishOptions
	Ack() error
}

type Message struct {
	Header map[string]string
	Body   []byte
}

func NewMessage() *Message {
	return &Message{
		Header: make(map[string]string),
	}
}

type Subscriber interface {
	Topic() string
	Options() SubscribeOptions
	Unsubscribe() error
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
