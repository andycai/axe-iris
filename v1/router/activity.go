package router

import (
	"axe/v1/system"

	"github.com/kataras/iris/v12"
)

func init() {
	routerNoCheck = append(routerNoCheck, registerActivityRouter)
}

func registerActivityRouter(app *iris.Application) {
	actsAPI := app.Party("/activities")
	{
		actsAPI.Get("/{aid:int64}", system.Activity.GetActivityById)
		actsAPI.Get("/", system.Activity.GetActivities)

		actsAPI.Post("/", system.Activity.Create)
		actsAPI.Post("/{aid:int64}/end", system.Activity.End)
		actsAPI.Post("/{aid:int64}/apply", system.Activity.Apply)
		actsAPI.Post("/{aid:int64}/cancel", system.Activity.Cancel)
		actsAPI.Post("/{aid:int64}/remove/:index", system.Activity.Remove)

		actsAPI.Put("/{aid:int64}", system.Activity.Update)
	}
}
