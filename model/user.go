package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Users struct {
	gorm.Model
	Id        uint
	Username  string
	Password  string
	Email     string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}
