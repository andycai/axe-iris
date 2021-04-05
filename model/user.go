package model

type User struct {
	Id         int64  `db:"id" json:"id"`
	Sex        int    `db:"sex" json:"sex"`
	Scores     int    `db:"scores" json:"scores"`
	Username   string `db:"username" json:"username"`
	Token      string `db:"token" json:"token"`
	WxToken    string `db:"wx_token" json:"wx_token"`
	WxNick     string `db:"wx_nick" json:"wx_nick"`
	Nick       string `db:"nick" json:"nick"`
	Ip         string `db:"ip" json:"ip"`
	Phone      string `db:"phone" json:"phone"`
	Email      string `db:"email" json:"email"`
	CreateAt   string `db:"create_at" json:"create_at"`
	Groups     string `db:"groups" json:"groups"`
	Activities string `db:"activities" json:"activities"`
}
