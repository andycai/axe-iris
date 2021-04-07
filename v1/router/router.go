package router

import "github.com/kataras/iris/v12/core/router"

var (
	routerNoCheck = make([]func(party router.Party), 0)
)
