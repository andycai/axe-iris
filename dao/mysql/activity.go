package mysql

type ActivityDao struct {
	//
}

var Activity = new(ActivityDao)

func (a ActivityDao) Create() {
	//
}

func (a ActivityDao) GetActivityById(id int64) {

}

func (a ActivityDao) GetActivitiesByType(typ, status, page, num int) {

}

func (a ActivityDao) GetActivitiesById(ids string) {

}

func (a ActivityDao) UpdateActivityStatus(id int64) {

}

func (a ActivityDao) UpdateActivityById(id int64) {

}
