package basetest

import (
	"fmt"
	"testing"
	"time"
)

func TestBingfa(t *testing.T) {
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

// 控制并发数，godoc， vfs 虚拟文件系统
// vfs/gatefs 控制访问该虚拟文件系统的最大并发数
func bingfa111() {
	/*
		fs := gatefs.New(vfs.OS("/path"), make(chan bool, 8))
		f :=fs.String()
		ff,_ := fs.Lstat("")
		ff.IsDir()
	*/
}

func TestChanSuijishu(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			select {
			case ch <- 0:
			case ch <- 1:
			case ch <- 2:
			case ch <- 3:
			}
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
