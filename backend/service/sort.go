package service

import (
	"bookManage/global"
	"bookManage/model"
)

var SortService = sortservice{}

type sortservice struct {
}

func (st sortservice) GetSort(pageSize int, offSetval int) (sortList []model.Sort, total int64) {
	global.DB.Model(&sortList).Count(&total).Limit(pageSize).Offset(offSetval).Find(&sortList)
	return
}

func (st sortservice) GetName(id int) string {
	sort := model.Sort{}
	global.DB.Where("id=?", id).Find(&sort)
	return sort.Name
}
