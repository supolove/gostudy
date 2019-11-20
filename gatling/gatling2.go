package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	fmt.Println("start server")
	http.HandleFunc("/account", handler2) // each request calls handler
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
