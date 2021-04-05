package system

import "github.com/kataras/iris/v12"

type GroupSystem struct{}

var Group = new(GroupSystem)

func (g GroupSystem) GetGroupById(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) GetGroupsByUserId(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) GetGroups(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) GetApplyList(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) GetActivitiesByGroupId(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) Create(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) Apply(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) Approve(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) Promote(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) Transfer(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) Remove(ctx iris.Context) {
	ctx.JSON(nil)
}

func (g GroupSystem) Quit(ctx iris.Context) {
	ctx.JSON(nil)
}

// put
func (g GroupSystem) Update(ctx iris.Context) {
	ctx.JSON(nil)
}
