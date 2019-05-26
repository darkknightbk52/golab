package registry

type Result struct {
	Action  string
	Service *Service
}

type Watcher interface {
	Next() (*Result, error)
	Stop()
}
