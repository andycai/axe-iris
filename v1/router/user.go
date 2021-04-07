package router

import (
	"axe/v1/system"
	"github.com/kataras/iris/v12"
)

func init() {
	routerNoCheck = append(routerNoCheck, registerUserRouter)
}

func registerUserRouter(app *iris.Application) {
	app.Post("/login", system.User.Login)
	app.Post("/login_wx", system.User.Login)
	app.Post("/login_wx", system.User.Logout)

	usersAPI := app.Party("/users")
	{
		usersAPI.Get("/{uid:int64}", system.User.GetUser)
		usersAPI.Get("/your/groups", system.Group.GetGroupsByUserId)
		usersAPI.Get("/your/activities", system.Activity.GetActivitiesByUserId)
	}
}
