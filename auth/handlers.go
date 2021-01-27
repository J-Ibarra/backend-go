package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginHandler(c *gin.Context) {
	payload := loginPayload{}
	c.BindJSON(&payload)

	user, err := FindUser(payload.Username)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid credentials",
			"error":   err.Error(),
		})
		return
	}

	err = ConfirmPassword(user, payload.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid credentials",
			"error":   err.Error(),
		})
		return
	}

	jwt, err := generateJwt(user.Username)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "could not generate token",
			"error":   err.Error(),
		})
		return
	}

	response := loginResponse{
		Token: jwt,
	}

	c.JSON(http.StatusOK, response)
}

func registerHandler(c *gin.Context) {
	payload := registerPayload{}
	c.BindJSON(&payload)

	_, err := FindUser(payload.Username)

	if err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "username has been take",
			"error":   "can not create user",
		})
		return
	}

	user := User{
		Username: payload.Username,
		Password: payload.Password,
	}

	CreateUser(user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "user create successfully",
		"user":    user,
	})

}

func getUserHandler(c *gin.Context) {
	user, err := FindUser(c.GetString("UserID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "could not find user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
