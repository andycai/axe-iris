package comp

import (
	"axe/define"
	"github.com/golang-module/carbon"
	"github.com/spf13/cast"
	"math"
)

type Activity struct {
	Id        int64   `json:"id"`
	Planner   int64   `json:"planner"`
	Kind      int     `json:"kind"`       // 活动分类:1羽毛球,2篮球,3足球,4聚餐...
	Type      int     `json:"type"`       // 活动类型:1全局保护,2全局公开,3群组
	Status    int     `json:"status"`     // 活动状态:1进行中,2正常结算完成,3手动终止
	Quota     int     `json:"quota"`      // 名额
	GroupId   int     `json:"group_id"`   // 群组ID
	Ahead     int     `json:"ahead"`      // 提前取消报名限制（小时）
	FeeType   int     `json:"fee_type"`   // 结算方式:1免费,2活动前,3活动后男女平均,4活动后男固定|女平摊,5活动后男平摊|女固定
	FeeMale   int     `json:"fee_male"`   // 男费用
	FeeFemale int     `json:"fee_female"` // 女费用
	Title     string  `json:"title"`
	Remark    string  `json:"remark"`
	Addr      string  `json:"addr"`
	BeginAt   string  `json:"begin_at"`
	EndAt     string  `json:"end_at"`
	Queue     []int64 `json:"queue"`     // 报名队列
	QueueSex  []int   `json:"queue_sex"` // 报名队列中的性别
}

func (a Activity) Init() {
	a.Queue = make([]int64, 0)
	a.QueueSex = make([]int, 0)
}

func (a Activity) InGroup() bool {
	return a.GroupId > 0
}

func (a Activity) IsPlanner(uid int64) bool {
	return uid == a.Planner
}

func (a Activity) Settle(fee int) {
	switch a.FeeType {
	case define.FeeTypeActivityAA:
		cost := math.Round(cast.ToFloat64(fee) / cast.ToFloat64(a.totalCount()))
		a.FeeMale = cast.ToInt(cost)
		a.FeeFemale = cast.ToInt(cost)
	case define.FeeTypeActivityAB:
		a.FeeFemale = cast.ToInt(math.Round(cast.ToFloat64(fee) - cast.ToFloat64(a.FeeMale*a.maleCount())))
	case define.FeeTypeActivityBA:
		a.FeeMale = cast.ToInt(math.Round(cast.ToFloat64(fee) - cast.ToFloat64(a.FeeFemale*a.femaleCount())))
	}
}

// 报名的人数超过候补的限制，避免乱报名，如带100000人报名
func (a Activity) OverQuota(total int) bool {
	return len(a.Queue)+total-a.Quota > define.ActivityOverFlow
}

// 要取消报名的数量超过已经报名的数量
func (a Activity) NotEnough(uid int64, total int) bool {
	c := 0
	for _, v := range a.Queue {
		if v == uid {
			c += 1
		}
	}
	return total > c
}

func (a Activity) InQueue(uid int64) bool {
	for _, v := range a.Queue {
		if v == uid {
			return true
		}
	}
	return false
}

func (a Activity) GetIdFromQueue(index int) int64 {
	if index < 0 || index >= len(a.Queue) {
		return 0
	}
	return a.Queue[index]
}

func (a Activity) Enqueue(uid int64, maleCount, femaleCount int) {
	a.fixQueue()
	for i := 0; i < maleCount; i++ {
		a.Queue = append(a.Queue, uid)
		a.QueueSex = append(a.QueueSex, define.TypeSexMale)
	}
	for i := 0; i < femaleCount; i++ {
		a.Queue = append(a.Queue, uid)
		a.QueueSex = append(a.QueueSex, define.TypeSexFemale)
	}
}

func (a Activity) Dequeue(index int) bool {
	a.fixQueue()
	if index < 0 || index >= len(a.Queue) {
		return false
	}
	a.Queue = append(a.Queue[:index], a.Queue[index+1:]...)
	a.QueueSex = append(a.QueueSex[:index], a.QueueSex[index+1:]...)
	return true
}

func (a Activity) DequeueMany(uid int64, maleCount, femaleCount int) {
	a.fixQueue()
	mCount := 0
	fCount := 0
	size := len(a.Queue)
	posArr := make([]int, 1)
	for i := size - 1; i >= 0; i-- {
		if a.Queue[i] == uid {
			// 男
			if a.QueueSex[i] == define.TypeSexMale && maleCount > mCount {
				mCount += 1
				posArr = append(posArr, i)
			}
			// 女
			if a.QueueSex[i] == define.TypeSexFemale && femaleCount > fCount {
				fCount += 1
				posArr = append(posArr, i)
			}
			if mCount >= maleCount && fCount >= femaleCount {
				break
			}
		}
	}
	for _, v := range posArr {
		a.Queue = append(a.Queue[:v], a.Queue[v+1:]...)
		a.QueueSex = append(a.QueueSex[:v], a.QueueSex[v+1:]...)
	}
}

// HasBegun 是否开始
func (a Activity) HasBegun() bool {
	return carbon.Parse(a.BeginAt).DiffInHours() <= 0
}

// CanCancel 能否取消报名
func (a Activity) CanCancel() bool {
	hours := carbon.Parse(a.BeginAt).DiffInHours()
	return cast.ToInt(hours) >= a.Ahead
}

// 私有方法

// 长度不一致，修正使其一致
func (a Activity) fixQueue() {
	df := len(a.QueueSex) - len(a.Queue)
	switch {
	case df > 0:
		a.QueueSex = a.QueueSex[:len(a.Queue)]
	case df < 0:
		a.Queue = a.Queue[:len(a.QueueSex)]
	}
}

// totalCount 有效的报名人数（不包括候选）
func (a Activity) totalCount() int {
	c := 0
	size := len(a.Queue)
	if a.Quota >= size {
		c = size
	} else {
		c = a.Quota
	}
	return c
}

func (a Activity) maleCount() int {
	c := 0
	total := a.totalCount()
	for i := 0; i < total; i++ {
		if len(a.QueueSex) > i && a.QueueSex[i] == define.TypeSexMale {
			c += 1
		}
	}
	return c
}

func (a Activity) femaleCount() int {
	c := 0
	total := a.totalCount()
	for i := 0; i < total; i++ {
		if len(a.QueueSex) > i && a.QueueSex[i] == define.TypeSexFemale {
			c += 1
		}
	}
	return c
}
