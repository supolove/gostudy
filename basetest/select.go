package basetest

import (
	"fmt"
	"runtime"
)

// case 语句里必须是一个io操作
func main5() {
	/*
		ch := make(chan int)	// 用来进行数据通信的channel
		quit := make(chan bool)	// 判断是否退出的channel

		go func() {
			for i:=0; i<5 ; i++  {
				ch <- i
				time.Sleep(time.Second)
			}
			close(ch)
			quit <- true
			runtime.Goexit()
		}()



		for{					// 主go程 读数据
			select {
			case num := <-ch:
				fmt.Println("读到：",num)
			case <- quit:
				// break			//跳出select
				return
			}
			fmt.Println("")
		}

	*/

	select2()
}

func fibonacci(ch chan int, quit chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Print(num, " ")
		case <-quit:
			runtime.Goexit()
		}
	}
}

// select的 斐波那契数列实现
func select2() {
	ch := make(chan int)
	quit := make(chan bool)

	go fibonacci(ch, quit)

	x, y := 1, 1
	for i := 0; i < 20; i++ {
		ch <- x
		x, y = y, x+y
	}

	quit <- true
}

// 超时
