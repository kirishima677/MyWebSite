package admin_services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"goMyWebSite/db"
	"goMyWebSite/model"
)

func GetArticlesListRowCountService() int64 {
	connection := db.Connection()
	defer connection.Close()

	var articles model.Articles
	var count int64
	connection.Model(&articles).Count(&count)

	return count
}

func GetArticlesListService(offset int, limit int) *gorm.DB {

	connection := db.Connection()
	defer connection.Close()

	var articles []model.Articles

	result := connection.Limit(limit).Offset(offset).Find(&articles).Group("id")

	fmt.Println("GetArticlesListService")

	return result
}

func UpsertArticlesListService() {

}

func var_dump(v ...interface{}) {
	for _, vv := range v {
		fmt.Printf("%#v\n", vv)
	}
}
