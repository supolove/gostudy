package basetest

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return ctx.Err()
		}
	}
}

func TestContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//context.WithCancel(context.Background())
	//context.WithValue()
	//context.WithDeadline()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	time.Sleep(2 * time.Second)
	cancel()

	wg.Wait()
}
