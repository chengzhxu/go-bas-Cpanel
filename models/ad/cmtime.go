package ad

/*
  @Author : czx
*/
import (
	"time"
)

type BaseModel struct {
	Ctime time.Time  `gorm:"column:ctime" json:"ctime" form:"ctime"`
	Mtime time.Time  `gorm:"column:mtime" json:"mtime" form:"mtime"`
}
