package system

import "github.com/kataras/iris/v12"

type GroupSystem struct {}

var group = new(GroupSystem)

func (g GroupSystem) getGroupById(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) getGroupsByUserId(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) getGroups(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) getApplyList(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) getActivitiesByGroupId(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) create(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) apply(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) approve(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) promote(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) transfer(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) remove(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) quit(ctx iris.Context) {
	ctx.JSON(nil)
}

// put
func (g GroupSystem) update(ctx iris.Context) {
	ctx.JSON(nil)
}