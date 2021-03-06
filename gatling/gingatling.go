package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "net/http/pprof"
)

// http://127.0.0.1:8000/account?sleep=10
func main() {
	fmt.Println("start server")
	//start := time.Now()

	//time.Sleep(time.Second)
	//fmt.Printf("%f", time.Since(start).Seconds())
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	//g := gin.Default()
	gin.DisableConsoleColor()

	g.POST("/account", handler3)
	g.Run(":8000")
}

// handler echoes the Path component of the request URL r.
func handler3(c *gin.Context) {
	c.Writer.Write([]byte{})
	//c.Writer.Write([]byte("ok"))
	//v := c.Query("sleep")
	//i, err := com.StrTo(v).Int()
	//if err != nil {
	//	panic(err)
	//}
	//time.Sleep(time.Second * time.Duration(i))
	//c.JSON(0, fmt.Sprintf("sleep=%d", i))
}
