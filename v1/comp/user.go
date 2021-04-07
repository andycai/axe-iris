package comp

import (
	"axe/util/slice"
)

type User struct {
	Id          int64   `json:"id"`
	Sex         int     `json:"sex"`
	Scores      int     `json:"scores"`
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	Token       string  `json:"token"`
	WxToken     string  `json:"wx_token" db:"wx_token"`
	WxNick      string  `json:"wx_nick" db:"wx_nick"`
	Nick        string  `json:"nick"`
	Ip          string  `json:"ip"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	CreateAt    string  `json:"create_at" db:"create_at"`
	Groups      string  `json:"-" db:"groups"`
	Activities  string  `json:"-" db:"activities"`
	GroupsV     []int   `json:"groups" db:"-"`
	ActivitiesV []int64 `json:"activities" db:"-"`
}

func NewUser() *User {
	u := new(User)
	u.GroupsV = make([]int, 0)
	u.ActivitiesV = make([]int64, 0)
	return u
}

func (u *User) Init() {
	json.Unmarshal([]byte(u.Groups), &u.GroupsV)
	json.Unmarshal([]byte(u.Activities), &u.ActivitiesV)
	//u.GroupsV = make([]int, 0)
	//u.ActivitiesV = make([]int64, 0)
}

func (u User) HasActivity(aid int64) bool {
	for _, activity := range u.ActivitiesV {
		if activity == aid {
			return true
		}
	}
	return false
}

func (u *User) AddActivity(aid int64) bool {
	if !u.HasActivity(aid) {
		u.ActivitiesV = append(u.ActivitiesV, aid)
		return true
	}
	return false
}

func (u *User) RemoveActivity(aid int64) bool {
	if u.HasActivity(aid) {
		u.ActivitiesV = slice.RemoveI64(u.ActivitiesV, aid)
		return true
	}
	return false
}

func (u User) HasGroup(gid int) bool {
	for _, group := range u.GroupsV {
		if group == gid {
			return true
		}
	}
	return false
}

func (u *User) AddGroup(gid int) bool {
	if !u.HasGroup(gid) {
		u.GroupsV = append(u.GroupsV, gid)
		return true
	}
	return false
}

func (u *User) RemoveGroup(gid int) bool {
	if u.HasGroup(gid) {
		u.GroupsV = slice.RemoveInt(u.GroupsV, gid)
		return true
	}
	return false
}
