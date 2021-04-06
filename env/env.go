package env

import "github.com/kataras/iris"

var app = iris.New()
var m = map[string]interface{}{

}

func GetApp() *iris.Application {
	return app
}

func GetMap()*map[string]interface{}{
	return &m
}



