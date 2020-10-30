package ad

import (
	"ferry/models/ad"
	"ferry/tools"
	"ferry/tools/app"
	"ferry/tools/app/msg"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/*
  @Author : x
*/

// @Summary 分页广告列表数据
// @Description 分页列表
// @Tags 广告
// @Param adid query int false "adid"
// @Param title query string false "title"
// @Param status query int false "status"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/adList [get]
// @Security
func GetAdList(c *gin.Context) {
	var (
		Ad   ad.Ad
		err    error
		result []ad.AdList
		pageSize  = 10
		pageIndex = 1
		count int
	)
	size := c.Request.FormValue("pageSize")
	if size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	index := c.Request.FormValue("pageIndex")
	if index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	Ad.AdId, _ = tools.StringToInt(c.Request.FormValue("adid"))
	Ad.Title = c.Request.FormValue("title")
	Ad.Status, _ = tools.StringToInt(c.Request.FormValue("status"))

	result, count, err = Ad.GetPage(true, pageSize, pageIndex)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	app.PageOK(c, result, count, pageSize, pageIndex, msg.GetSuccess)
}


// @Summary 获取指定广告数据
// @Description 获取广告
// @Tags 广告
// @Param adid path string false "adid"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/ad/{adid} [get]
// @Security
func GetAd(c *gin.Context) {
	var (
		err  error
		Ad ad.Ad
	)
	Ad.AdId, _ = tools.StringToInt(c.Param("adid"))

	result, err := Ad.Get()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, result, msg.GetSuccess)
}

// @Summary 创建广告
// @Description 获取JSON
// @Tags 广告
// @Accept  application/json
// @Product application/json
// @Param data body ad.Ad true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/ad [post]
// @Security Bearer
func CreateAd(c *gin.Context) {
	var data ad.AdRes
	err := c.BindWith(&data, binding.JSON)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	data.UserId, _ = tools.StringToInt(tools.GetUserIdStr(c))
	result, err := data.Create()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, result, msg.CreatedSuccess)
}

// @Summary 修改广告
// @Description 获取JSON
// @Tags 广告
// @Accept  application/json
// @Product application/json
// @Param adid path int true "adid"
// @Param data body ad.Ad true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/ad [put]
// @Security Bearer
func UpdateAd(c *gin.Context) {
	var data ad.Ad
	err := c.BindJSON(&data)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	result, err := data.Update(data.AdId)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, result, msg.UpdatedSuccess)
}


