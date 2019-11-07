package basetest

import (
	"fmt"
	"testing"
	"time"
)

/*
3种定时方法
1. time.Sleep(time.Second)
2. Timer.C
3. time.After
*/

func TestDingshiqi(t *testing.T) {

	fmt.Println("当前时间：", time.Now())

	// 创建定时器
	//myTimer := time.NewTimer(time.Second * 2)
	//nowTime := <- myTimer.C
	//fmt.Println("现在时间：", nowTime)
	//
	//nowTim := <- time.After(time.Second * 2)
	//fmt.Println("现在时间", nowTim)

	// timer2()

	// 时间戳
	// int32(time.Now().Unix())
	fmt.Println(time.Now().Unix())

}

// 定时器的停止和重置
func timer1() {
	mytimer := time.NewTimer(time.Second * 3)
	go func() {
		<-mytimer.C
		fmt.Println("子go程序，定时完毕")
	}()
	//mytimer.Stop()
	for {

	}
}

// 定时器周期定时
func timer2() {
	myTicker := time.NewTicker(time.Second)
	go func() {
		for {
			nowTime := <-myTicker.C
			fmt.Println("nowTime: ", nowTime)
		}
	}()

	for {

	}
}
