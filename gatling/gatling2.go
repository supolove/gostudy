package main

import (
	"fmt"
	"github.com/Unknwon/com"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

// 本地拉取代码
func main() {
	port := ""
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	if port == "" {
		fmt.Println("请输入端口号")
		return
	}
	fmt.Println("端口号:", port)
	fmt.Println("start server")
	http.HandleFunc("/pull", handler2) // each request calls handler
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

// handler echoes the Path component of the request URL r.
func handler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	fmt.Println("start pull content")
	out, e, err := com.ExecCmd("git", "pull")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(out)
	fmt.Println(e)
}
