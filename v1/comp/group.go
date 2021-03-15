package comp

import "axe/define"

type Group struct {
	Id         int64
	Level      int
	Scores     int
	Name       string
	Logo       string
	Notice     string
	Addr       string
	Activities []int64
	Pending    []int64 // 申请入群列表
	Members    []Member
}

func (g Group) Init() {
	g.Activities = make([]int64, 0)
	g.Pending = make([]int64, 0)
	g.Members = make([]Member, 0)
}

func (g Group) NotInPending(index int) bool {
	return index < 0 || index >= len(g.Pending)
}

// IsMember 是否群成员
func (g Group) IsMember(uid int64) bool {
	for _, member := range g.Members {
		if member.Id == uid {
			return true
		}
	}
	return false
}

// IsOwner 是否群主
func (g Group) IsOwner(uid int64) bool {
	for _, member := range g.Members {
		if member.Id == uid && member.Pos == define.PositionGroupOwner {
			return true
		}
	}
	return false
}

// IsManager 是否管理员
func (g Group) IsManager(uid int64) bool {
	for _, member := range g.Members {
		if member.Id == uid && member.Pos == define.PositionGroupManager {
			return true
		}
	}
	return false
}

func (g Group) ManagerCount() int {
	c := 0
	for _, member := range g.Members {
		if member.Pos > define.PositionGroupMember {
			c += 1
		}
	}
	return c
}

func (g *Group) ExistActivity(aid int64) bool {
	for _, v := range g.Activities {
		if v == aid {
			return true
		}
	}
	return false
}

func (g Group) AddActivity(aid int64) {
	if !g.ExistActivity(aid) {
		g.Activities = append(g.Activities, aid)
	}
}

func (g *Group) Promote(uid int64) bool {
	for _, member := range g.Members {
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
	for _, member := range g.Members {
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
	for i, member := range g.Members {
		if member.Id == uid {
			index = i
			break
		}
	}
	if index >= 0 {
		g.Members = append(g.Members[:index], g.Members[index+1:]...)
		return true
	}
	return false
}

func (g Group) NotIn(uid int64) bool {
	for _, member := range g.Members {
		if member.Id == uid {
			return false
		}
	}
	return true
}
