package ad

import (
	"encoding/json"
	"ferry/global/orm"
	"time"
	_ "time"
)

/*
  @Author : lanyulei
*/

type Ad struct {
	AdId   int    `json:"adid" gorm:"primary_key;AUTO_INCREMENT;column:adid"` //广告id
	Title string    `json:"title" gorm:"type:varchar(512);"`            //广告标题
	TimeStart time.Time `json:"time_start" gorm:"type:timestamp;column:time_start"`     //广告开始时间
	TimeEnd time.Time `json:"time_end" gorm:"type:timestamp;column:time_end"`       //广告结束时间
	UserId   int    `juseridson:"" gorm:"type:int(11);column:userid"`              //创建用户
	Status   int    `json:"status" gorm:"type:tinyint(1);"`              //状态
	ReConf     json.RawMessage `gorm:"column:re_conf; type:json" json:"re_conf" form:"re_conf"`    //配置
	FormatId    int `json:"formatid" gorm:"type:int(1);column:formatid"`           //广告形式
	Priority   int `json:"priority" gorm:"type:int(11);column:priority"`               //优先级
	MgtvVar     json.RawMessage `gorm:"column:mgtv_var; type:json" json:"mgtv_var" form:"mgtv_var"`    //MGTV 参数配置
	SupSwitch int `json:"sup_switch" gorm:"type:tinyint(1);column:sup_switch"`   //矩阵-光芒广告状态
	SupStatus int `json:"sup_status" gorm:"type:tinyint(1);column:sup_status"`   //辅助渠道启停状态
	MonitorType int `json:"monitor_type" gorm:"type:tinyint(1);column:monitor_type"`   //监测方式
	BasMonitorType int `json:"bas_monitor_type" gorm:"type:tinyint(1);column:bas_monitor_type"`   //BAS 监测方式
	IsPromote int `json:"is_promote" gorm:"type:tinyint(1);column:is_promote"`   //是否推广广告
	BaseModel
}

type AdSup struct {
	ReApp     json.RawMessage `gorm:"column:re_app; type:json" json:"re_app" form:"re_app"`    //MGTV 参数配置
	SupConf     json.RawMessage `gorm:"column:sup_conf; type:json" json:"sup_conf" form:"sup_conf"`    //矩阵 参数配置
	SupApp     json.RawMessage `gorm:"column:sup_app; type:json" json:"sup_app" form:"sup_app"`    //辅助渠道
}

type AdSchedule struct {
	Schedule     json.RawMessage `gorm:"column:schedule; type:json" json:"schedule" form:"schedule"`    //排期
}

type AdWorktime struct {
	WorktimeSwitch int `json:"worktime_switch"  gorm:"type:tinyint(1);column:worktime_switch"`      //投放时段 是否全天 0：是  1：否
	Worktime     json.RawMessage `gorm:"column:worktime; type:json" json:"worktime" form:"worktime"`    //投放时段
}

type AdRes struct {
	Ad
	AdSup
	AdSchedule
	AdWorktime
}


func (Ad) TableName() string {
	return "ads_ad"
}

type AdGroup struct {
	GroupId int `gorm:"type:int(11);column:groupid" json:"groupid"`
	GroupTitle string `gorm:"type:varchar(512);column:group_title" json:"group_title"`
}

type AdLable struct {
	AdId       int         `gorm:"-" json:"adid"`
	Title    string      `gorm:"-" json:"title"`
}


type AdList struct {
	Ad
	AdGroup
}

func (e *Ad) Create() (Ad, error) {
	var ad Ad
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return ad, err
	}
	ad = *e
	return ad, nil
}

func (e *Ad) Get() (Ad, error) {
	var ad Ad

	table := orm.Eloquent.Table(e.TableName())
	if e.AdId != 0 {
		table = table.Where("adid = ?", e.AdId)
	}
	if e.Title != "" {
		table = table.Where("title = ?", e.Title)
	}

	if err := table.First(&ad).Error; err != nil {
		return ad, err
	}
	return ad, nil
}

func (e *Ad) GetList() ([]Ad, error) {
	var ad []Ad

	table := orm.Eloquent.Table(e.TableName())
	if e.AdId != 0 {
		table = table.Where("adid = ?", e.AdId)
	}
	if e.Title != "" {
		table = table.Where("title = ?", e.Title)
	}
	if e.Status > -1 {
		table = table.Where("status = ?", e.Status)
	}

	if err := table.Order("ctime desc").Find(&ad).Error; err != nil {
		return ad, err
	}
	return ad, nil
}

func (e *Ad) GetPage(bl bool, pageSize int, pageIndex int) ([]AdList, int, error) {
	var adList []AdList
	//orm.Eloquent.LogMode(true)

	table := orm.Eloquent.Table(e.TableName() + " as ad ").Select("ad.*, g.groupid, g.title as group_title").
	Joins("inner JOIN ads_ad_group as ag on ag.adid = ad.adid").
	Joins("inner JOIN ads_group as g on g.groupid = ag.groupid")
	if e.AdId != 0 {
		table = table.Where("ad.adid = ?", e.AdId)
	}
	if e.Title != "" {
		table = table.Where("ad.title like ?", "%"+e.Title+"%")
	}
	if e.Status > -1 {
		table = table.Where("ad.status = ?", e.Status)
	}

	var count int

	if err := table.Order("ad.ctime desc").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&adList).Error; err != nil {
		return nil, 0, err
	}
	table.Count(&count)
	return adList, count, nil
}

func (e *Ad) Update(id int) (update Ad, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("adid = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据

	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}

	return
}

