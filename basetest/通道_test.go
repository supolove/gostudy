package basetest

import (
	"fmt"
	"log"
	"testing"
	"time"
)

/*


 */

var channel = make(chan int, 10)

func printer(s string) {
	for _, ch := range s {
		c := fmt.Sprintf("%c", ch)
		fmt.Println(c)
		time.Sleep(300 * time.Millisecond)
	}

}

func person1() {
	//printer("hello")
	for i := 0; i < 5; i++ {
		channel <- i
		fmt.Println("写:", i)
	}

}

func person2() {
	//num := <- channel
	for {
		i := <-channel
		//fmt.Println("读:",i)
		log.Print("读:", i, " ")
		log.Print(len(channel))
		//printer("world")
	}

}

func TestTongdao(t *testing.T) {
	// channel 队列模型
	// 是一种数据类型，对应一个"管道"（通道）	<-channel
	//	printer("world")
	// 解决协程的同步问题以及协程之间数据共享（数据传递）的问题

	// 写端  ch <- "hello" 	写端写数据，读端不在读，阻塞
	// 读端  temp <- ch 		读端读数据，同时写端不在写， 读端阻塞

	fmt.Println("456")

	go person2()
	//go person1()

	for {
		time.Sleep(time.Second)
		if len(channel) == 0 {
			for i := 0; i < 5; i++ {
				channel <- i
			}
		}
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
