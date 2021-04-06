package middleware

import (
	"fmt"
	"github.com/kataras/iris"
)

func PageMiddleware(ctx iris.Context){

	ctx.Values().Set("keu", "name")
	fmt.Println(ctx.Values().GetString("keu"))

	ctx.Next()
}

func Recover(ctx iris.Context) {
	defer func() {
		if err := recover() ;err!= nil{
			fmt.Println(err)
			ctx.JSON(map[string]interface{}{
				"code" : -1,
				"msg": err,
			})
		}
	}()

	ctx.Next()
}
