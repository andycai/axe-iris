package system

import "github.com/kataras/iris/v12"

type ActivitySystem struct {}

var ACtivity = new(ActivitySystem)

func (a ActivitySystem) getActivityById(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) getActivitiesByUserId(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) getActivities(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) create(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) end(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) apply(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) cancel(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) remove(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) update(ctx iris.Context) {
	ctx.JSON(nil)
}