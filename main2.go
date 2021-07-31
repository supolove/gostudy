package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func test() {
	sig := make(chan os.Signal)
	stopCh := make(chan struct{})
	finishedCh := make(chan struct{})

	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	go func(stopCh, finishedCh chan struct{}) {
		for {
			select {
			case <-stopCh:
				fmt.Println("stopped")
				finishedCh <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(stopCh, finishedCh)

	<-sig
	stopCh <- struct{}{}
	<-finishedCh
	fmt.Println("finished")
}

func test2() {
	sig := make(chan os.Signal)
	stopCh := make(chan chan struct{})

	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	go func(stopCh chan chan struct{}) {
		for {
			select {
			case ch := <-stopCh:
				fmt.Println("stopped")
				ch <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(stopCh)

	<-sig
	ch := make(chan struct{})
	stopCh <- ch
	<-ch
	fmt.Println("finished")
}

func test3() {
	sig := make(chan os.Signal)
	ctx, cancel := context.WithCancel(context.Background())
	finishedCh := make(chan struct{})

	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	go func(ctx context.Context, finishedCh chan struct{}) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopped")
				finishedCh <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(ctx, finishedCh)

	<-sig
	cancel()
	<-finishedCh
	fmt.Println("finished")
}
