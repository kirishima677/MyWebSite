package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Articles struct {
	gorm.Model
	Id           uint `gorm:"AUTO_INCREMENT"`
	Title        string
	Body         string
	PostedUserId uint
	ReleaseFlg   bool
	//CreatedAt    time.Time
	//UpdatedAt    time.Time
	//DeletedAt    *time.Time
}
