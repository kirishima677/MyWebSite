package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Users struct {
	gorm.Model
	Id       uint `gorm:"AUTO_INCREMENT"`
	LoginId  string
	Password string
	Email    string
	//CreatedAt time.Time
	//UpdatedAt time.Time
	//DeletedAt time.Time
}
