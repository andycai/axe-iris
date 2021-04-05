package comp

type Session struct {
	At    int64  `json:"at"`
	Uid   int64  `json:"uid"`
	Sex   int    `json:"sex"`
	Token string `json:"token"`
}
