package basetest

import (
	"fmt"
	"testing"
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

func TestPrime2(t *testing.T) {
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

/*-----------------------------------------------*/

func generateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func primeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			var i = <-in
			fmt.Println("计算：", i, prime)
			if i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func TestPrime(t *testing.T) {
	ch := generateNatural()

	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = primeFilter(ch, prime)
	}
}
