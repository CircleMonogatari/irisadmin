package api

import (
	"fmt"
	"github.com/kataras/iris"
	"irisadmin/env"
)

func Init()  {
	fmt.Println("加载 api router")
	app := env.GetApp()
	api := app.Party("/api")
	api.Get("/", func(ctx iris.Context) {

	})
}
