package registry

type Registry interface {
	Init(...Option) error
	Options() Options
	Register(*Service, ...RegisterOption) error
	Deregister(*Service) error
	GetService(string) ([]*Service, error)
	ListService() ([]*Service, error)
	Watch(...WatchOption) (Watcher, error)
	String() string
}

type Option func(*Options)

type WatchOption func(*WatchOptions)

type RegisterOption func(*RegisterOptions)
