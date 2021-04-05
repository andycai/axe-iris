package comp

type Player struct {
	Id     int64  `json:"id"`
	Sex    int    `json:"sex"`
	WxNick string `json:"wx_nick"`
	Nick   string `json:"nick"`
}
