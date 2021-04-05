package comp

type Member struct {
	Id     int64  `json:"id"`
	Scores int    `json:"scores"`
	Pos    int    `json:"pos"`
	Sex    int    `json:"sex"`
	At     int64  `json:"at"`
	WxNick string `json:"wx_nick"`
	Nick   string `json:"nick"`
}
