package registry

import "context"

type Options struct {
	Context context.Context
}

type Option func(*Options)

type Node struct {
	ID       string
	Address  string
	Port     string
	MetaData map[string]string
}

type Value struct {
	Name   string
	Type   string
	Values []*Value
}

type Endpoint struct {
	Name     string
	Request  *Value
	Response *Value
	MetaData map[string]string
}

type Service struct {
	Name      string
	Version   string
	MetaData  map[string]string
	Endpoints []*Endpoint
	Nodes     []*Node
}

type WatchOptions struct {
}

type WatchOption func(*WatchOptions)

type Result struct {
	Action  string
	Service *Service
}

type Watcher interface {
	Next() (*Result, error)
	Stop() error
}

type RegisterOptions struct {
}

type RegisterOption func(*RegisterOptions)

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
