package basetest

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := make(chan int, 1)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子go程 i = ", i)
			ch <- i
		}
	}()

	time.Sleep(time.Second * 2)

	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("主go程读:", num)
	}
}
