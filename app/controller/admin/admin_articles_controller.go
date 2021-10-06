package admin_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goMyWebSite/db"
	"goMyWebSite/model"
	"goMyWebSite/services/admin"
	"net/http"
	"strconv"
)

func AdminArticlesControllerIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_articles_index.tmpl", gin.H{
		"title": "記事管理 ダッシュボード",
	})
	fmt.Println("/admin/articles")
}

type paginationInfo struct {
	RowCount int
	Offset   int
	Limit    int
	Path     string
}

func AdminArticlesControllerList(c *gin.Context) {

	var offset, _ = strconv.Atoi(c.Query("offset"))
	var limit, _ = strconv.Atoi(c.Query("limit"))

	if limit <= 0 {
		limit = 5
	}

	rowCount := admin_services.GetArticlesListRowCountService()
	result := admin_services.GetArticlesListService(offset, limit)

	pagination := paginationInfo{}
	pagination.RowCount = int(rowCount)
	pagination.Offset = offset
	pagination.Limit = limit
	pagination.Path = "/admin/articles/list"

	c.HTML(http.StatusOK, "admin_articles_list.tmpl", gin.H{
		"title":      "記事の一覧",
		"list":       result.Value,
		"pagination": pagination,
	})

	fmt.Println("/admin/articles/list")
}

func AdminArticlesControllerNew(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_articles_new.tmpl", gin.H{
		"title": "記事の投稿",
	})
	fmt.Println("/admin/articles/new")
}

func AdminArticlesControllerNewPost(c *gin.Context) {
	fmt.Println("POST")
	fmt.Println("/admin/articles/new")

	// @todo バリデーションを追加する

	// フォームから値を受け取る
	releaseFlg := c.PostForm("release_flg")
	title := c.PostForm("title")
	body := c.PostForm("body")
	fmt.Println(releaseFlg)
	fmt.Println(title)
	fmt.Println(body)

	// 登録処理
	connection := db.Connection()
	defer connection.Close()

	var articles model.Articles

	//t := time.Now()

	articles.Title = c.PostForm("title")
	articles.Body = c.PostForm("body")
	articles.PostedUserId = 1
	articles.ReleaseFlg, _ = strconv.ParseBool(c.PostForm("release_flg"))
	// 日付はGorm ORMがよしなにやってくれる
	//articles.CreatedAt = t
	//articles.UpdatedAt = t

	result := connection.Create(&articles)
	//result := connection.First(&articles, 1).Related(&articles.Id)
	fmt.Println(result)

	// @todo エラーハンドリング

	c.Redirect(http.StatusSeeOther, "/admin/articles/list")
}

func AdminArticlesControllerEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_articles_edit.tmpl", gin.H{
		"title": "記事の編集",
	})
	fmt.Println("/admin/articles/edit")
}

func AdminArticlesControllerDelete(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_articles_index.tmpl", gin.H{
		"title": "記事管理 ダッシュボード",
	})
	fmt.Println("/admin/articles")
}
