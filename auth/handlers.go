package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginHandler(c *gin.Context) {
	payload := loginPayload{}
	c.BindJSON(&payload)

	jwt := generateJwt(payload.Username)

	response := loginResponse{
		Token: jwt,
	}

	c.JSON(http.StatusOK, response)
}

func getUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"user": c.GetString("UserID"),
	})
}
