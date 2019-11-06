package basetest

import (
	"time"
)

func main9() {
	// 并行（parallel）：指在同一时刻，cpu执行多条指令
	// 并发 用户体验上，程序在并行执行，多个计划任务，轮询执行
	// 程序： 编译成功得到的二进制文件，只占用磁盘空间
	// 进程： 运行起来的程序，占用系统资源（内存 主要是 堆内存）

	// 协程（coroutine）：轻量级的线程

	// 并发手段 gorotine，channel

	// runtime go运行误操作

	// 用来出让当前go程所占用的时间片
	// runtime.Gosched()

	// 立即终止gorotine执行
	// runtime.Goexit()

	// 设置cpu核心数，1为单核
	// runtime.GOMAXPROCS(1)

	// runtime.GC()

	go func() {
		time.Sleep(2)
	}()
}
