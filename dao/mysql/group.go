package mysql

type GroupDao struct {
}

var Group = new(GroupDao)

func (g *GroupDao) Create() {
	//
}

func (g GroupDao) GetGroupById(id int) {

}

func (g GroupDao) GetGroups(page, num int) {

}

func (g GroupDao) GetGroupsByIds(ids string) {

}

func (g GroupDao) UpdateGroupById(id int) {

}
