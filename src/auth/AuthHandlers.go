package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c *gin.Context) {
	payload := loginPayload{}
	c.BindJSON(&payload)

	response := map[string]interface{}{
		"username": payload.Username,
		"password": payload.Password,
	}

	c.JSON(http.StatusOK, response)
}

// ConfigureRoutes configure routes auth
func ConfigureRoutes(router *gin.Engine) {
	auth := router.Group("/api/v1/auth")

	auth.POST("/sign-in", login)
}
