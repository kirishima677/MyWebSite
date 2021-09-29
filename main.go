package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	_ "go.uber.org/zap"
	"goMyWebSite/db"
	"goMyWebSite/middleware"
	"goMyWebSite/model"
	"net/http"
)

func main() {
	// DBから取得するサンプル
	connection := db.Connection()
	defer connection.Close()

	var user model.Users
	result := connection.First(&user, 1).Related(&user.Id)
	fmt.Println("##### connection result #####")
	fmt.Println(result.Value)
	fmt.Println("##### connection result #####")

	// zapロガーサンプル
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
	router.Use(middleware.CommonMiddleware())

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
	router.GET("/file", middleware.AuthCheckMiddleware(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "file.tmpl", gin.H{
			"title": "file list page",
		})
		fmt.Println("/file")
	})

	// file upload
	router.GET("/file/upload", middleware.AuthCheckMiddleware(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "file_upload.tmpl", gin.H{
			"title": "file upload page",
		})
		fmt.Println("/file/upload")
	})

	// zapでのログの出力サンプル
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	logger.Debug("debug")
	logger.Info("info", zap.String("key", "value"))

	sugar = logger.Sugar()
	sugar.Warn("warning sugar")
	sugar.Error("error sugar")

	//起動とサーバーポートの指定
	router.Run(":8080")
}
