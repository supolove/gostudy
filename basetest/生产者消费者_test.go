package basetest

import (
	"fmt"
	"testing"
	"time"
)

/**
生产者消费者
生产者（发送数据端）	->	缓冲区	->	消费者（接收数据端）
缓冲区好处
1.解耦（降低生产者和消费者之间耦合度）
2.提高并发能力
3.缓存(生产者和消费者数据处理速度不一致，暂存数据)
*/

// 生产者
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
		fmt.Println("生产：", i*i)
	}
	close(out)
}

// 消费者
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("消费者拿到:", num)
		time.Sleep(1)
	}
}

// 模拟订单
type OrderInfo struct {
	id int
}

func producer2(out chan<- OrderInfo) {
	for i := 0; i < 10; i++ {
		order := OrderInfo{id: i + 1}
		out <- order
	}
	close(out)
}

func consumer2(in <-chan OrderInfo) {
	for order := range in {
		fmt.Println(order.id)
	}
}

// 消费者

func TestProducer(t *testing.T) {
	ch := make(chan int, 5)

	go producer(ch)

	consumer(ch)

}
