package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ConfigureRoutes configure routes auth
func ConfigureRoutes(router *gin.Engine) {
	auth := router.Group("/api/v1/auth")

	auth.POST("/sign-in", loginHandler)

	secureGroup := auth.Group("")
	secureGroup.Use(authMiddleware())

	secureGroup.GET("/user", getUserHandler)
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeaders := c.Request.Header["Authorization"]

		if len(authHeaders) < 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Authorization",
			})
		}

		for _, authorization := range authHeaders {
			auth := DecodeToken{}
			if err := VerifyJwt(authorization, &auth); err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Invalid Authorization",
				})
				break
			}
			c.Set("UserID", string(auth.Id))
			c.Next()
			break
		}
	}
}
