package comp

import (
	"axe/define"
)

type Group struct {
	Id          int64    `json:"id"`
	Level       int      `json:"level"`
	Scores      int      `json:"scores"`
	Name        string   `json:"name"`
	Logo        string   `json:"logo"`
	Notice      string   `json:"notice"`
	Addr        string   `json:"addr"`
	Activities  string   `json:"-"`
	Pending     string   `json:"-"`
	Members     string   `json:"-"`
	ActivitiesV []int64  `json:"activities"`
	PendingV    []int64  `json:"pending"` // 申请入群列表
	MembersV    []Member `json:"members"`
}

func NewGroup() *Group {
	g := new(Group)
	return g
}

func (g Group) Init() {
	json.Unmarshal([]byte(g.Activities), &g.ActivitiesV)
	json.Unmarshal([]byte(g.Pending), &g.PendingV)
	json.Unmarshal([]byte(g.Members), &g.MembersV)
	//g.ActivitiesV = make([]int64, 0)
	//g.PendingV = make([]int64, 0)
	//g.MembersV = make([]Member, 0)
}

func (g Group) NotInPendingV(index int) bool {
	return index < 0 || index >= len(g.PendingV)
}

// IsMember 是否群成员
func (g Group) IsMember(uid int64) bool {
	for _, member := range g.MembersV {
		if member.Id == uid {
			return true
		}
	}
	return false
}

// IsOwner 是否群主
func (g Group) IsOwner(uid int64) bool {
	for _, member := range g.MembersV {
		if member.Id == uid && member.Pos == define.PositionGroupOwner {
			return true
		}
	}
	return false
}

// IsManager 是否管理员
func (g Group) IsManager(uid int64) bool {
	for _, member := range g.MembersV {
		if member.Id == uid && member.Pos == define.PositionGroupManager {
			return true
		}
	}
	return false
}

func (g Group) ManagerCount() int {
	c := 0
	for _, member := range g.MembersV {
		if member.Pos > define.PositionGroupMember {
			c += 1
		}
	}
	return c
}

func (g *Group) ExistActivity(aid int64) bool {
	for _, v := range g.ActivitiesV {
		if v == aid {
			return true
		}
	}
	return false
}

func (g Group) AddActivity(aid int64) {
	if !g.ExistActivity(aid) {
		g.ActivitiesV = append(g.ActivitiesV, aid)
	}
}

func (g *Group) Promote(uid int64) bool {
	for _, member := range g.MembersV {
		if member.Id == uid {
			member.Pos = define.PositionGroupManager
			return true
		}
	}
	return false
}

func (g *Group) Transfer(uid, mid int64) bool {
	b := false
	if !g.IsMember(uid) || !g.IsMember(mid) {
		return false
	}
	for _, member := range g.MembersV {
		// 外部自行判断权限
		if member.Id == uid {
			member.Pos = define.PositionGroupMember
		}
		if member.Id == mid {
			member.Pos = define.PositionGroupOwner
			b = true
		}
	}
	return b
}

func (g *Group) Remove(uid int64) bool {
	index := -1
	for i, member := range g.MembersV {
		if member.Id == uid {
			index = i
			break
		}
	}
	if index >= 0 {
		g.MembersV = append(g.MembersV[:index], g.MembersV[index+1:]...)
		return true
	}
	return false
}

func (g Group) NotIn(uid int64) bool {
	for _, member := range g.MembersV {
		if member.Id == uid {
			return false
		}
	}
	return true
}
