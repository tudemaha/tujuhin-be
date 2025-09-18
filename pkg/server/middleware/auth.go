package middleware

import "github.com/gin-gonic/gin"

type AuthMiddleware interface {
	Auth() gin.HandlerFunc
}
