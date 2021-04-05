package model

type Group struct {
	Id         int64  `db:"id" json:"id"`
	Level      int    `db:"level" json:"level"`
	Scores     int    `db:"scores" json:"scores"`
	Name       string `db:"name" json:"name"`
	Logo       string `db:"logo" json:"logo"`
	Notice     string `db:"notice" json:"notice"`
	Addr       string `db:"addr" json:"addr"`
	Activities string `db:"activities" json:"activities"`
	Pending    string `db:"pending" json:"pending"` // 申请入群列表
	Members    string `db:"members" json:"members"`
}
