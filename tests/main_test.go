package main

import (
	_ "bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	_ "mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	// パラメータ生成
	var param gin.Param = gin.Param{Key: "classificationId", Value: "1"}
	var params gin.Params = gin.Params{param}

	//fmt.Println(params)

	//var router = router()
	var w *httptest.ResponseRecorder = httptest.NewRecorder()

	// リクエスト生成
	req, _ := http.NewRequest("GET", "/", nil)
	//router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// Contextセット
	var context *gin.Context = &gin.Context{Request: req, Params: params}

	huga(context)
}

func huga(c *gin.Context) {

}
