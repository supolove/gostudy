package basetest

import (
	"fmt"
	"testing"
	"time"
)

/*


 */

var channel = make(chan int)

func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}

}

func person1() {
	printer("hello")
	channel <- 8
}

func person2() {
	//num := <- channel
	<-channel
	printer("world")
}

func TestTongdao(t *testing.T) {
	// channel 队列模型
	// 是一种数据类型，对应一个"管道"（通道）
	// 解决协程的同步问题以及协程之间数据共享（数据传递）的问题

	// 写端  ch <- "hello" 	写端写数据，读端不在读，阻塞
	// 读端  temp <- ch 		读端读数据，同时写端不在写， 读端阻塞
	go person1()
	go person2()

	for {

	}

	/*
		关闭close(ch)
		数据不发送完，不应该关闭
		已经关闭的channel，不应该写数据
		已经关闭的channel，可以读到数据
		写段已经关闭channel，可以从中读取数据，读到0，说明，写端关闭
	*/
}

func send(out chan<- int) {
	out <- 89
	close(out)
}

func recv(in <-chan int) {
	n := <-in
	fmt.Println("读取:", n)
}

func channelSingle() {
	ch := make(chan int)

	go func() {
		send(ch)
	}()
	recv(ch)
}
