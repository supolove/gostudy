package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {

	beego.Post("/account", func(ctx *context.Context) {
		ctx.Output.Body([]byte{})
		//ctx.Output.JSON("",false, false)
	})

	beego.Get("/account", func(ctx *context.Context) {
		ctx.Output.Body([]byte{})
	})

	beego.Run(":8000")
}
