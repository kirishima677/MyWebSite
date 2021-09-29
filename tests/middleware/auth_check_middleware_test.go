package middleware

import (
	"goMyWebSite/middleware"
	"testing"
)

func TestAuthCheckMiddleware(t *testing.T) {
	middleware.AuthCheckMiddleware()
}
