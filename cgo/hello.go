package main

import "C"
import (
	"fmt"
	"time"
)

const End = 50

func source(ch chan<- int) {
	for i := 2; i < End; i++ {
		ch <- i
	}
}

func validate(in <-chan int, out chan<- int, fix int) {
	for {
		select {
		case i := <-in:
			fmt.Println("i", i)
			if i%fix != 0 {
				out <- i
			}
		case <-time.After(1 * time.Second):
			close(out)
			return
		}
	}
}

func main() {
	// 逐个获取待检测的数据源
	ch := make(chan int)
	go source(ch)

	for {
		data, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(data)
		out := make(chan int)
		go validate(ch, out, data)
		ch = out
	}
}
