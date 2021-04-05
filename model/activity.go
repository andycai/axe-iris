package model

type Activity struct {
	Id        int64  `json:"id"`
	Planner   int64  `json:"planner"`
	Kind      int    `json:"kind"`                       // 活动分类:1羽毛球,2篮球,3足球,4聚餐...
	Type      int    `json:"type"`                       // 活动类型:1全局保护,2全局公开,3群组
	Status    int    `json:"status"`                     // 活动状态:1进行中,2正常结算完成,3手动终止
	Quota     int    `json:"quota"`                      // 名额
	GroupId   int    `json:"group_id" db:"group_id"`     // 群组ID
	Ahead     int    `json:"ahead"`                      // 提前取消报名限制（小时）
	FeeType   int    `json:"fee_type" db:"fee_type"`     // 结算方式:1免费,2活动前,3活动后男女平均,4活动后男固定|女平摊,5活动后男平摊|女固定
	FeeMale   int    `json:"fee_male" db:"fee_male"`     // 男费用
	FeeFemale int    `json:"fee_female" db:"fee_female"` // 女费用
	Title     string `json:"title"`
	Remark    string `json:"remark"`
	Addr      string `json:"addr"`
	BeginAt   string `json:"begin_at" db:"begin_at"`
	EndAt     string `json:"end_at" db:"end_at"`
	Queue     string `json:"queue"`                    // 报名队列
	QueueSex  string `json:"queue_sex" db:"queue_sex"` // 报名队列中的性别
}
