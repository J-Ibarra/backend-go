package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginHandler(c *gin.Context) {
	payload := loginPayload{}
	c.BindJSON(&payload)

	response := map[string]interface{}{
		"username": payload.Username,
		"password": payload.Password,
	}

	c.JSON(http.StatusOK, response)
}
