package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func CommonMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before logic")
		c.Next()
		log.Println("after logic")
	}
}
