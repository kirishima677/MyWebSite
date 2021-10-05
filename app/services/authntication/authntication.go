package authntication

import (
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	//"log"
	_ "net/http"
)

func Logout(c *gin.Context) {

	//セッションからデータを破棄する
	//var session = Default(c)
	//log.Println("セッション取得")
	//session.Clear()
	//log.Println("クリア処理")
	//if result := session.Save(); result.Error != nil {
	//	// エラーハンドリング
	//	errorHandling()
	//}
}

func Login(c *gin.Context, UserId string) {

	//セッションにデータを格納する
	//session := Default(c)
	//
	//session.Set("UserId", UserId)
	//session.Save()
	//if result := session.Save(); result.Error != nil {
	//	// エラーハンドリング
	//	errorHandling()
	//}
}

func errorHandling() {

}
