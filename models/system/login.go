package system

import (
	"errors"
	"ferry/global/orm"
	"ferry/tools"
	"github.com/haxqer/gintools/logging"
)

/*
  @Author : lanyulei
*/

type Login struct {
	Username  string `form:"UserName" json:"username" binding:"required"`
	Password  string `form:"Password" json:"password" binding:"required"`
	//LoginType int    `form:"LoginType" json:"loginType"`
}

func (u *Login) GetUser() (user SysUser, role SysRole, e error) {

	e = orm.Eloquent.Table("user").Where("username = ? ", u.Username).First(&user).Error

	if e != nil {
		logging.Error(e)
		return user, role, e
	}

	// 验证密码
		//_, e = tools.CompareHashAndPassword(user.Password, u.Password)
		p := u.Password + user.Salt
		r := tools.CompareMD5AndPassword(user.Password, p)
		if !r {
			e = errors.New("password mistake")
			return
		}

	e = orm.Eloquent.Table("user_role").Where("id = ? ", user.RoleId).First(&role).Error
	if e != nil {
		logging.Error(e)
		return
	}
	return
}
