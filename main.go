package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
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
