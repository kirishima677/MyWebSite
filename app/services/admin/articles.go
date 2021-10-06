package admin_services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"goMyWebSite/db"
	"goMyWebSite/model"
)

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
