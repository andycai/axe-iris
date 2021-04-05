package cache

import "axe/v1/comp"

type GroupCache struct {
	groups map[int]*comp.Group
}

var Group = &GroupCache{
	groups: make(map[int]*comp.Group),
}
