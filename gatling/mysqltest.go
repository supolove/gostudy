package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"sync"
	"time"
)

func main() {

	var engine *xorm.Engine

	type Account struct {
		Id      int64  `xorm:"not null pk autoincr INT(11)"`
		Name    string `xorm:"unique"`
		Balance float64
	}

	fmt.Println("start server")

	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@tcp(192.168.11.150:3306)/gatling?charset=utf8")
	if err != nil {
		panic(err)
	}

	engine.SetMaxOpenConns(40)

	if err = engine.Sync2(new(Account)); err != nil {
		panic(err)
	}

	start := time.Now()
	/*
		for i := 0; i < 5; i++ {
			go func(v int) {
				for j := 0; j < 100000; j++ {
					name := fmt.Sprintf("name%d%d", v, j)
					_, err = engine.Insert(Account{Name: name, Balance: 20.0})
					if err != nil {
						panic(err)
					}
					fmt.Println("insert!!:", name)
				}
				fmt.Printf("%d-%f\n", v, time.Since(start).Seconds())
			}(i)
		}
	*/

	wg := sync.WaitGroup{}
	for j := 0; j < 500000; j++ {
		wg.Add(1)
		go func(v int) {
			name := fmt.Sprintf("name%d", v)
			_, err = engine.Insert(Account{Name: name, Balance: 20.0})
			if err != nil {
				panic(err)
			}
			fmt.Println("insert!!:", name)
			wg.Done()
		}(j)
		if j%10 == 0 {
			wg.Wait()
		}
	}

	fmt.Printf("%f\n", time.Since(start).Seconds())

	//fmt.Printf("%f\n", time.Since(start).Seconds())
	for {

	}
}
