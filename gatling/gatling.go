package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

var engine *xorm.Engine
var mychan chan Message

type Account struct {
	Id      int64  `xorm:"not null pk autoincr INT(11)"`
	Name    string `xorm:"unique"`
	Balance float64
}

type Message struct {
	Name   string
	Writer http.ResponseWriter
}

func main() {
	fmt.Println("start server")
	mychan = make(chan Message, 100)

	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/gatling?charset=utf8")
	if err != nil {
		panic(err)
	}

	if err = engine.Sync2(new(Account)); err != nil {
		panic(err)
	}

	engine.SetMaxOpenConns(20)

	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:10000", nil))
	}()

	http.HandleFunc("/account", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))

}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	//m := Message{Name:r.FormValue("cmbUid"), Writer:w}
	//mychan <- r.FormValue("cmbUid")
	//mychan <- m
	w.Write([]byte("ok"))

	fmt.Println(runtime.NumGoroutine())

}
