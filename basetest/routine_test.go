package basetest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func p() {
	for i := 0; i < 2; i++ {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func sell(c chan int) {
	for {
		num := <-c
		fmt.Println("sell", num, "head")
	}
}

func produce(c chan int) {
	for {
		num := rand.Intn(10)
		t := time.Duration(num)
		fmt.Println("product", num, "bread")
		c <- num
		time.Sleep(time.Second * t)
	}
}

func TestRoutine(t *testing.T) {
	go p()
	var input string
	fmt.Scanln(&input)
	fmt.Println("End")

	//go routine使用channel来进行routine间的通信
	//显示结果也正面，channel的使用规则
	var c chan int = make(chan int)
	go sell(c)
	go produce(c)
	var input2 string
	fmt.Scanln(&input2)
	fmt.Println("end")
}
