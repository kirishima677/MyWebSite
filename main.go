package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	_ "go.uber.org/zap"
	"goMyWebSite/db"
	"goMyWebSite/middleware"
	"goMyWebSite/model"
	"net/http"
)

// Connection
func Connection() redis.Conn {
	const Addr = "127.0.0.1:6379"

	c, err := redis.Dial("tcp", Addr)
	if err != nil {
		panic(err)
	}
	return c
}

// データの登録(Redis: SET key value)
func Set(key, value string, c redis.Conn) string {
	res, err := redis.String(c.Do("SET", key, value))
	if err != nil {
		panic(err)
	}
	return res
}

// データの取得(Redis: GET key)
func Get(key string, c redis.Conn) string {
	res, err := redis.String(c.Do("GET", key))
	if err != nil {
		panic(err)
	}
	return res
}

func main() {
	// Redis接続
	c := Connection()
	defer c.Close()

	// データの登録(Redis: SET key value)
	res_set := Set("sample-key", "sample-value", c)
	fmt.Println(res_set) // OK

	// データの取得(Redis: GET key)
	res_get := Get("sample-key", c)
	fmt.Println(res_get) // sample-value

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
