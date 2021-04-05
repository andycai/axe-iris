package cache

import "axe/v1/comp"

type ActivityCache struct {
	activities map[int64]*comp.Activity
}

var Activity = &ActivityCache{
	activities: make(map[int64]*comp.Activity),
}
