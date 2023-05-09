package controller

import (
	"bookManage/schemas/resp"
	"bookManage/service"
	"bookManage/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

var SortApi = sortApi{}

type sortApi struct {
}

func (sa sortApi) GetSortList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	offSetval := util.ToolsUtil.Paging(pageSize, pageNum)
	list, total := service.SortService.GetSort(pageSize, offSetval)
	data := gin.H{
		"total":    total,
		"pageNum":  offSetval,
		"pageSize": pageSize,
		"list":     list,
	}
	if len(list) == 0 {
		resp.OkWithMsg(c, "没有查询到数据")
	} else {
		resp.OkWithMsgData(c, "查询成功", data)
	}
}

func (sa sortApi) GetSortName(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	name := service.SortService.GetName(id)
	resp.OkWithMsgData(c, "查询成功", name)
}
