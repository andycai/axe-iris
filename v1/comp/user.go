package comp

import (
	"axe/model"
	"axe/util/slice"
)

type User struct {
	Id         int64   `json:"id"`
	Sex        int     `json:"sex"`
	Scores     int     `json:"scores"`
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	Token      string  `json:"token"`
	WxToken    string  `json:"wx_token" db:"wx_token"`
	WxNick     string  `json:"wx_nick" db:"wx_nick"`
	Nick       string  `json:"nick"`
	Ip         string  `json:"ip"`
	Phone      string  `json:"phone"`
	Email      string  `json:"email"`
	CreateAt   string  `json:"create_at" db:"create_at"`
	Groups     []int   `json:"groups"`
	Activities []int64 `json:"activities"`
}

func NewUser(user *model.User) *User {
	u := new(User)
	u.Init(user)
	return u
}

func (u *User) Init(user *model.User) {
	//u.groups = make([]int, 0)
	//u.Activities = make([]int64, 0)
	u.Id = user.Id
	u.Sex = user.Sex
	u.Scores = user.Scores
	u.Username = user.Username
	u.Token = user.Token
	u.WxToken = user.WxToken
	u.WxNick = user.WxNick
	u.Nick = user.Nick
	u.Ip = user.Ip
	u.Phone = user.Phone
	u.Email = user.Email
	u.CreateAt = user.CreateAt
	json.Unmarshal([]byte(user.Groups), &u.Groups)
	json.Unmarshal([]byte(user.Activities), &u.Activities)
}

func (u *User) ToModel() *model.User {
	user := new(model.User)
	user.Id = u.Id
	user.Sex = u.Sex
	user.Scores = u.Scores
	user.Username = u.Username
	user.Token = u.Token
	user.WxToken = u.WxToken
	user.WxNick = u.WxNick
	user.Nick = u.Nick
	user.Ip = u.Ip
	user.Phone = u.Phone
	user.Email = u.Email
	user.CreateAt = u.CreateAt
	groups, _ := json.Marshal(&u.Groups)
	user.Groups = string(groups)
	activities, _ := json.Marshal(&u.Activities)
	user.Activities = string(activities)

	return user
}

func (u User) HasActivity(aid int64) bool {
	for _, activity := range u.Activities {
		if activity == aid {
			return true
		}
	}
	return false
}

func (u *User) AddActivity(aid int64) bool {
	if !u.HasActivity(aid) {
		u.Activities = append(u.Activities, aid)
		return true
	}
	return false
}

func (u *User) RemoveActivity(aid int64) bool {
	if u.HasActivity(aid) {
		u.Activities = slice.RemoveI64(u.Activities, aid)
		return true
	}
	return false
}

func (u User) HasGroup(gid int) bool {
	for _, group := range u.Groups {
		if group == gid {
			return true
		}
	}
	return false
}

func (u *User) AddGroup(gid int) bool {
	if !u.HasGroup(gid) {
		u.Groups = append(u.Groups, gid)
		return true
	}
	return false
}

func (u *User) RemoveGroup(gid int) bool {
	if u.HasGroup(gid) {
		u.Groups = slice.RemoveInt(u.Groups, gid)
		return true
	}
	return false
}
