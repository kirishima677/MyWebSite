package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func AuthCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before authCheckMiddleware")
		c.Next()
		log.Println("after authCheckMiddleware")
	}
}
