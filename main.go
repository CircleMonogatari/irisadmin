package main

import (
	"fmt"
	"github.com/kataras/iris"
	"irisadmin/env"
	"irisadmin/middleware"
	"irisadmin/route"
	_ "irisadmin/route/api"
)

func main() {
	fmt.Println("go go go")

	app := env.GetApp()
	(*env.GetMap())["qqq"] = 1

	// 捕获异常
	app.Use(middleware.Recover)

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>demo</h1>")
	})

	app.Handle("GET", "/panic", func(ctx iris.Context) {
		panic("测试异常")
	})

	// 初始化配置文件
	env.ConfigInit("config.json")
	// 注册路由
	route.Init()

	app.Run(iris.Addr(":8002"))

}
