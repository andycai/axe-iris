package define

const (
	TypeActivityNone = iota
	TypeActivityProtected // 全局不公开活动
	TypeActivityPublic // 全局公开活动
	TypeActivityGroup // 群组活动
)

const (
	StatusActivityNone = iota
	StatusActivityDoing
	StatusActivityDone
	StatusActivityEnd
)

const (
	KindActivityNone = iota
	KindActivityDate // 约饭
	KindActivityBadminton // 羽毛球
)

const (
	FeeTypeActivityNone = iota
	FeeTypeActivityFree // 免费
	FeeTypeActivityBefore // 活动前结算
	FeeTypeActivityAA // 全部分摊
	FeeTypeActivityAB // 男固定，女平摊
	FeeTypeActivityBA // 男平摊，女固定
)

// 候补列表不能超过固定数量
const ActivityOverFlow = 10
