package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	_ "go.uber.org/zap"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Debug("debug")
	logger.Info("info", zap.String("key", "value"))

	sugar := logger.Sugar()
	sugar.Warn("warning sugar")
	sugar.Error("error sugar")

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	// ミドルウエアを使う
	router.Use(commonMiddleware())

	// index
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// login
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title": "login page",
		})
		fmt.Println("/login")
	})
	router.POST("/login", func(c *gin.Context) {
		// 認証処理

		// 認証できなかったらログインフォームの表示
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title":   "login page",
			"message": "",
		})

		fmt.Println("/login POST")
	})

	// file list
	router.GET("/file", authCheckMiddleware(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "file.tmpl", gin.H{
			"title": "file list page",
		})
		fmt.Println("/file")
	})

	// file upload
	router.GET("/file/upload", authCheckMiddleware(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "file_upload.tmpl", gin.H{
			"title": "file upload page",
		})
		fmt.Println("/file/upload")
	})

	//起動とサーバーポートの指定
	router.Run(":8080")
}

func commonMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before logic")
		c.Next()
		log.Println("after logic")
	}
}

func authCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before authCheckMiddleware")
		c.Next()
		log.Println("after authCheckMiddleware")
	}
}

//func Connection() *gorm.DB {
//	db, err := gorm.Open("mysql", "root:@tcp(db:3306)/gin_app?charset=utf8&parseTime=True&loc=Local")
//	if err != nil {
//		panic("failed to connect database")
//	}
//	db.LogMode(true)
//	return db
//}
