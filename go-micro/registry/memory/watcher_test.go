package memory

import (
	"github.com/darkknightbk52/golab/go-micro/registry"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func Test(t *testing.T) {
	RegisterTestingT(t)

	expectedResults := []*registry.Result{
		{
			Service: &registry.Service{
				Name: "service",
			},
			Action: "Insert",
		},
		{
			Service: &registry.Service{
				Name: "service",
			},
			Action: "Update",
		},
	}

	sentResults := []*registry.Result{
		{
			Service: &registry.Service{
				Name: "otherService",
			},
			Action: "Delete",
		},
	}
	sentResults = append(sentResults, expectedResults...)

	w := NewWatcher(registry.WatchService("service"))
	var notifiedResults []*registry.Result
	exit := make(chan bool)
	go func() {
		for {
			select {
			case <-exit:
				return
			default:
				r, _ := w.Next()
				notifiedResults = append(notifiedResults, r)
			}
		}
	}()

	for _, r := range expectedResults {
		w.res <- r
	}

	time.Sleep(time.Second)
	close(exit)

	t.Log("Compare expected & notified results")
	Expect(len(expectedResults)).Should(Equal(len(notifiedResults)), "difference between expected & notified results")

	for _, er := range expectedResults {
		seen := false
		for _, ne := range notifiedResults {
			if ne.Action == er.Action {
				seen = true
				Expect(ne.Service.Name).Should(Equal(er.Service.Name), "action not from watched service: "+ne.Service.Name)
			}
		}
		Expect(seen).Should(BeTrue())
		t.Log(er, er.Service.Name)
	}
}
