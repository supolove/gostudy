package basetest

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 死锁

/*
单go程自己死锁
go程间channel访问顺序导致死锁
多go程，多channel交叉死锁
*/
func TestSuo(t *testing.T) {
	sisuo3()
}

// 1.同一个goroutine中，使用同一个 channel 读写。
func sisuo1() {
	ch := make(chan int)
	ch <- 789
	num := <-ch //  这里会发生一直阻塞的情况，执行不到下面一句
	fmt.Println("num = ", num)
}

// 2.读写channel 先于go程创建
func sisuo2() {

	ch := make(chan int)
	num := <-ch
	fmt.Println("num = ", num)

	go func() {
		ch <- 789
	}()
}

// 3.2个以上的go程中，使用多个 channel 通信。 A go 程 获取channel 1 的同时，尝试使用channel 2，
// 同一时刻，B go 程 获取channel 2 的同时，尝试使用channel 1
// 是互相等对方造成死锁
func sisuo3() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println(num)
				//ch2 <- num
			}
		}
	}()
	ch1 <- 1

	for {
		select {
		case num := <-ch2:
			//fmt.Println(num)
			ch1 <- num
		}
	}
}

//

// 使用channel完成同步
var ch = make(chan int)

func printer2(str string) {
	for _, ch := range str {
		fmt.Println("c%", ch)
		time.Sleep(time.Millisecond * 300)
	}
}

func person3() { //先执行
	printer2("hello")
	ch <- 98
}

func person4() { //后执行
	<-ch
	printer2("world")
}

func fucisuo() {
	go person3()
	go person4()
}

//使用锁完成同步
// lock, unlock
var mutex sync.Mutex

func printer3(str string) {
	mutex.Lock()
	for _, ch := range str {
		fmt.Println("c%", ch)
		time.Sleep(time.Millisecond * 300)
	}
	mutex.Unlock()
}
func fucisuo2() {

}

// 读写锁
