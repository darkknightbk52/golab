package memory

import (
	"github.com/darkknightbk52/golab/go-micro/registry"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Watcher struct {
	id   string
	wo   registry.WatchOptions
	res  chan *registry.Result
	exit chan bool
}

func NewWatcher(opts ...registry.WatchOption) *Watcher {
	options := registry.WatchOptions{}

	for _, opt := range opts {
		opt(&options)
	}

	return &Watcher{
		id:   uuid.New().String(),
		wo:   options,
		res:  make(chan *registry.Result),
		exit: make(chan bool),
	}
}

func (w *Watcher) Next() (*registry.Result, error) {
	for {
		select {
		case r := <-w.res:
			if len(w.wo.Service) > 0 && w.wo.Service != r.Service.Name {
				continue
			}
			return r, nil
		case <-w.exit:
			return nil, errors.New("watcher stopped")
		}
	}
}

func (w *Watcher) Stop() {
	select {
	case <-w.exit:
	default:
		close(w.exit)
	}
}
