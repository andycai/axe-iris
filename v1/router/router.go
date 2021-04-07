package router

import (
	"github.com/kataras/iris/v12"
)

var (
	routerNoCheck = make([]func(*iris.Application), 0)
)

func InitRouter(app *iris.Application) {
	for _, f := range routerNoCheck {
		f(app)
	}
}
