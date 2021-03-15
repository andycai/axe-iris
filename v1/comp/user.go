package comp

import "axe/util/slice"

type User struct {
	Id         int64
	Sex        int
	Scores     int
	Username   string
	Password   string
	Token      string
	WxToken    string
	WxNick     string
	Nick       string
	Ip         string
	Phone      string
	Email      string
	CreateAt   string
	groups     []int
	activities []int64
}

func (u *User) Init() {
	u.groups = make([]int, 0)
	u.activities = make([]int64, 0)
}

func (u User) HasActivity(aid int64) bool {
	for _, activity := range u.activities {
		if activity == aid {
			return true
		}
	}
	return false
}

func (u *User) AddActivity(aid int64) bool {
	if !u.HasActivity(aid) {
		u.activities = append(u.activities, aid)
		return true
	}
	return false
}

func (u *User) RemoveActivity(aid int64) bool {
	if u.HasActivity(aid) {
		u.activities = slice.RemoveI64(u.activities, aid)
		return true
	}
	return false
}

func (u User) HasGroup(gid int) bool {
	for _, group := range u.groups {
		if group == gid {
			return true
		}
	}
	return false
}

func (u *User) AddGroup(gid int) bool {
	if !u.HasGroup(gid) {
		u.groups = append(u.groups, gid)
		return true
	}
	return false
}

func (u *User) RemoveGroup(gid int) bool {
	if u.HasGroup(gid) {
		u.groups = slice.RemoveInt(u.groups, gid)
		return true
	}
	return false
}
