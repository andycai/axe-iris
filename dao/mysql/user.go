package mysql

import (
	"axe/gl"
	"axe/model"
	"fmt"
)

type UserDao struct {
	//
}

var User = new(UserDao)

func (u UserDao) Create() {
	//fields := "username,password,token,nick,wx_token,wx_nick,sex,phone,email,ip,activities,groups"
	//sql := "INSERT INTO user (?) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"

	//result, err := db.Exec(sql, $fields, )
}

func (u *UserDao) GetUserByName(name string) {
	//
}

func (u *UserDao) GetUserById(id int64) *model.User {
	fields := "id,scores,username,token,nick,wx_token,wx_nick,sex,phone,email,ip,activities,groups,create_at"
	sql := fmt.Sprintf("SELECT %s FROM `user` WHERE id = ?", fields)

	user := new(model.User)
	if err := db.Get(user, sql, id); err != nil {
		gl.App.Logger().Errorf("Get user data failed, err: %v\n", err)
	}

	return user
}

func (u *UserDao) GetUsersByIds(ids string) []model.User {
	fields := "id,scores,username,token,nick,wx_token,wx_nick,sex,phone,email,ip,activities,groups,create_at"
	sql := fmt.Sprintf("SELECT %s FROM `user` WHERE id in(?)", fields)

	var users []model.User
	if err := db.Select(&users, sql, ids); err != nil {
		gl.App.Logger().Errorf("Get user data failed, err: %v\n", err)
	}

	return users
}

func (u UserDao) UpdateUserById(id int64) {

}
