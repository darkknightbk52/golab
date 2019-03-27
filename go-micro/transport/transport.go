package transport

type Options struct {
}

type Option func(*Options)

type DialOptions struct {
}

type DialOption func(*DialOptions)

type Message struct {
	Header map[string]string
	Body   []byte
}

type Socket interface {
	Send(*Message) error
	Recv(*Message) error
	Close() error
	Local() string
	Remote() string
}

type Client interface {
	Socket
}

type ListenOptions struct {
}

type ListenOption func(*ListenOptions)

type Listener interface {
	Accept(func(Socket)) error
	Close() error
	Addr() string
}

type Transport interface {
	Init(...Option) error
	Options() Options
	Dial(addr string, opts ...DialOption) (Client, error)
	Listen(addr string, opts ...ListenOption) (Listener, error)
	String() string
}
