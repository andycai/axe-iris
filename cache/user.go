package cache

import (
	"axe/dao/mysql"
	"axe/gl"
	"axe/v1/comp"
)

type UserCache struct {
	usersForName map[string]*comp.User
	usersForId   map[int64]*comp.User
	sessions     map[string]*comp.Session
}

var User = &UserCache{
	usersForName: make(map[string]*comp.User),
	usersForId:   make(map[int64]*comp.User),
	sessions:     make(map[string]*comp.Session),
}

func (uc *UserCache) GetUserById(id int64) *comp.User {
	if user, ok := uc.usersForId[id]; ok {
		gl.App.Logger().Infof("获得用户数据（缓存）：%d", id)
		return user
	}
	gl.App.Logger().Infof("获得用户数据（DB）：%d", id)
	user := mysql.User.GetUserById(id)
	user.OutDB()
	uc.usersForId[id] = user

	return user
}
