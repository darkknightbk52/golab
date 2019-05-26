package registry

type Node struct {
	Id       string
	Address  string
	Port     int
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
