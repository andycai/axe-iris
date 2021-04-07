package system

import (
	"axe/cache"
	"github.com/kataras/iris/v12"
)

type UserSystem struct{}

var User = new(UserSystem)

func (u UserSystem) GetUser(c iris.Context) {
	uid, _ := c.Params().GetInt64("uid")
	user := cache.User.GetUserById(uid)
	Ok(c, user)
}

func (u UserSystem) Login(c iris.Context) {
	c.JSON(nil)
}

func (u UserSystem) Logout(c iris.Context) {
	c.JSON(nil)
}

func (u UserSystem) EnterGroup(c iris.Context) {
	c.JSON(nil)
}

func (u UserSystem) QuitGroup(c iris.Context) {
	c.JSON(nil)
}

func (u UserSystem) ApplyActivity(c iris.Context) {
	c.JSON(nil)
}

func (u UserSystem) Cancel(c iris.Context) {
	c.JSON(nil)
}

func (u UserSystem) SaveData(c iris.Context) {
	c.JSON(nil)
}
