package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "net/http/pprof"
	"time"
)

func main() {
	fmt.Println("start server")
	start := time.Now()

	time.Sleep(time.Second)
	fmt.Printf("%f", time.Since(start).Seconds())

	//g := gin.New()
	g := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	g.GET("/account", handler3)
	g.Run(":8000")
}

// handler echoes the Path component of the request URL r.
func handler3(c *gin.Context) {
	c.JSON(0, "ok")
}
