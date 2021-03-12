package system

import "github.com/kataras/iris/v12"

type UserSystem struct {}

var User = new(UserSystem)

func (u UserSystem) getUser(ctx iris.Context) {
	ctx.JSON(nil)
}

func (u UserSystem) login(ctx iris.Context) {
	ctx.JSON(nil)
}

func (u UserSystem) logout(ctx iris.Context) {
	ctx.JSON(nil)
}

func (u UserSystem) enterGroup(ctx iris.Context) {
	ctx.JSON(nil)
}

func (u UserSystem) quitGroup(ctx iris.Context) {
	ctx.JSON(nil)
}

func (u UserSystem) applyActivity(ctx iris.Context) {
	ctx.JSON(nil)
}

func (u UserSystem) cancel(ctx iris.Context) {
	ctx.JSON(nil)
}

func (u UserSystem) saveData(ctx iris.Context) {
	ctx.JSON(nil)
}
