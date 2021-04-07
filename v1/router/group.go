package router

import (
	"axe/v1/system"
	"github.com/kataras/iris/v12"
)

func init() {
	routerNoCheck = append(routerNoCheck, registerGroupRouter)
}

func registerGroupRouter(app *iris.Application) {
	groupsAPI := app.Party("/groups")
	{
		groupsAPI.Get("/{gid:int}", system.Group.GetGroupById)
		groupsAPI.Get("/", system.Group.GetGroups)
		groupsAPI.Get("/{gid:int}/pending", system.Group.GetApplyList)
		groupsAPI.Get("/{gid:int}/activities", system.Group.GetActivitiesByGroupId)

		groupsAPI.Post("/", system.Group.Create)
		groupsAPI.Post("/{gid:int}/gl.Apply", system.Group.Apply)
		groupsAPI.Post("/{gid:int}/gl.Approve", system.Group.Approve)
		groupsAPI.Post("/{gid:int}/promote/:mid", system.Group.Promote)
		groupsAPI.Post("/{gid:int}/transfer/:mid", system.Group.Transfer)
		groupsAPI.Post("/{gid:int}/remove/:mid", system.Group.Remove)
		groupsAPI.Post("/{gid:int}/quit", system.Group.Quit)

		groupsAPI.Put("/{gid:int}", system.Group.Update)
	}
}
