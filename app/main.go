package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	_ "go.uber.org/zap"
	"goMyWebSite/controller/admin"
	"goMyWebSite/controller/user"
	"goMyWebSite/db"
	"goMyWebSite/middleware"
	"goMyWebSite/model"
	"goMyWebSite/redis"

	"goMyWebSite/services/authntication"
	"net/http"
)

func main() {
	// Redis接続
	c := redis.Connection()
	defer c.Close()

	// データの登録(Redis: SET key value)
	resSet := redis.Set("sample-key", "sample-value", c)
	fmt.Println(resSet) // OK

	// データの取得(Redis: GET key)
	resGet := redis.Get("sample-key", c)
	fmt.Println(resGet) // sample-value

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

	// zapでのログの出力サンプル
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	logger.Debug("debug")
	logger.Info("info", zap.String("key", "value"))

	sugar = logger.Sugar()
	sugar.Warn("warning sugar")
	sugar.Error("error sugar")

	router := gin.Default()

	// セッションCookieの設定
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*/*")
	//router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLGlob("templates/layout/*")

	// ミドルウエアを使う
	router.Use(middleware.CommonMiddleware())

	/**
	 * ユーザー系
	 */

	// index
	router.GET("/", user_controller.Index)

	// login
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title": "login page",
		})
		fmt.Println("/login")
	})
	router.POST("/login", func(c *gin.Context) {
		// 認証処理

		loginId := c.PostForm("id")
		password := c.PostForm("password")
		fmt.Println(loginId)
		fmt.Println(password)

		var user model.Users
		if result := connection.Where("login_id = ? and password = ?", loginId, password).First(&user); result.Error != nil {
			// ここでエラーハンドリング
			fmt.Println("error")
			// 認証できなかったらログインフォームの表示
			c.HTML(http.StatusOK, "view/login.tmpl", gin.H{
				"title":   "login page",
				"message": "",
			})
		} else {
			// レコードが存在した
			fmt.Println("##### connection result #####")
			fmt.Println(result.Value)
			fmt.Println("##### connection result #####")
			//
			var random, _ = MakeRandomStr(10)
			fmt.Println(random)
			//
			////セッションにデータを格納する
			authntication.Login(c, random)
			//authntication.Login(c, random)
			//
			fmt.Println("Login")
			c.Redirect(http.StatusSeeOther, "/")
		}

		fmt.Println("/login POST")
	})

	// logout
	router.GET("/logout", func(c *gin.Context) {
		authntication.Logout(c)
		c.HTML(http.StatusOK, "templates/view/admin_articles_index.tmpl", gin.H{
			"title": "Main website11",
		})
		fmt.Println("/logout")
	})

	/**
	 * 管理系
	 */
	// 管理 入口
	router.GET("/admin", admin_controller.AdminIndex)
	// 記事管理 ダッシュボード
	router.GET("/admin/articles", admin_controller.AdminArticlesControllerIndex)
	// 記事の一覧
	router.GET("/admin/articles/list", admin_controller.AdminArticlesControllerList)
	// 記事の投稿
	router.GET("/admin/articles/new", admin_controller.AdminArticlesControllerNew)
	router.POST("/admin/articles/new", admin_controller.AdminArticlesControllerNewPost)
	// 記事の編集
	router.GET("/admin/articles/edit", admin_controller.AdminArticlesControllerEdit)
	router.POST("/admin/articles/edit", admin_controller.AdminArticlesControllerEdit)
	// 記事の削除
	router.DELETE("/admin/articles/delete", admin_controller.AdminArticlesControllerDelete)

	//起動とサーバーポートの指定
	router.Run(":3000")
}

func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
