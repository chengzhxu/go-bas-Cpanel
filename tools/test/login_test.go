package test

import (
	"ferry/tools"
	"fmt"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	//var user system.SysUser
	//fmt.Printf("%v", 312313)
	//e := orm.Eloquent.Table("user").Where("username = ? ", "kevin").Find(&user).Error
	p :=  "1234560.812497550534321"
	mp := "843e0cb57a041b9840ad24833899d4ae"
	//str := tools.md5V(p)
	fmt.Printf("%v\n", p)

	r := tools.CompareMD5AndPassword(mp, p)

	if !r {
		fmt.Printf("%v\n", "FALSEEEEEEE")
	}

	//fmt.Printf("%v", user)

	os.Exit(0)
}
