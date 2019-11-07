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

func sisuo1() {
	ch := make(chan int)
	ch <- 789
	num := <-ch
	fmt.Println("num = ", num)
}

func sisuo2() {
	ch := make(chan int)
	num := <-ch
	fmt.Println("num = ", num)

	go func() {
		ch <- 789
	}()
}

func sisuo3() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case num := <-ch1:
				ch2 <- num
			}
		}
	}()

	for {
		select {
		case num := <-ch2:
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
	for {

	}
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
