package auth

import "github.com/gin-gonic/gin"

// ConfigureRoutes configure routes auth
func ConfigureRoutes(router *gin.Engine) {
	auth := router.Group("/api/v1/auth")

	auth.POST("/sign-in", loginHandler)
}
