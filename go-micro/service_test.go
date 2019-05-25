package micro

import (
	"context"
	"fmt"
	"github.com/darkknightbk52/golab/go-micro/client"
	"github.com/darkknightbk52/golab/go-micro/registry/memory"
	"log"
	"sync"
	"testing"
)

func TestService(t *testing.T) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	service := testService(ctx, &wg, "test.service")

	go func() {
		wg.Wait()

		err := testRequest(service.Client(), "test.service")
		if err != nil {
			t.Fatal(err)
		}

		testShutdown(&wg, cancel)
	}()

	err := service.Run()
	if err != nil {
		t.Fatal(err)
	}
}

func testShutdown(wg *sync.WaitGroup, cancel context.CancelFunc) {
	wg.Add(1)
	cancel()
	wg.Wait()
}

func testRequest(c client.Client, name string) error {
	req := c.NewRequest(
		name,
		"Debug.Health",
		nil,
	)

	var rsp interface {
		Status() string
	}
	err := c.Call(context.TODO(), req, rsp)
	if err != nil {
		return err
	}

	if rsp.Status() != "ok" {
		return fmt.Errorf("error response: %v", rsp)
	}

	log.Println(rsp)
	return nil
}

func testService(ctx context.Context, wg *sync.WaitGroup, name string) Service {
	wg.Add(1)

	registry := memory.NewRegistry()
	registry.(*memory.Registry).Setup()
	return NewService(
		Name(name),
		Context(ctx),
		Registry(registry),
		AfterStart(func() error {
			wg.Done()
			return nil
		}),
		AfterStop(func() error {
			wg.Done()
			return nil
		}),
	)
}
