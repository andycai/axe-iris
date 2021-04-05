package system

import "github.com/kataras/iris/v12"

type ActivitySystem struct{}

var Activity = new(ActivitySystem)

func (a ActivitySystem) GetActivityById(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) GetActivitiesByUserId(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) GetActivities(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) Create(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) End(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) Apply(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) Cancel(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) Remove(ctx iris.Context) {
	ctx.JSON(nil)
}

func (a ActivitySystem) Update(ctx iris.Context) {
	ctx.JSON(nil)
}
