package middleware

import (
	"goMyWebSite/app/middleware"
	"testing"
)

func TestAuthCheckMiddleware(t *testing.T) {
	middleware.AuthCheckMiddleware()
}
