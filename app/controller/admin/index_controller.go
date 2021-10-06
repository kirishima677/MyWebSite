package admin_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_index.tmpl", gin.H{
		"title": "管理 入口",
	})
	fmt.Println("/admin")
}
