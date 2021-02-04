package auth

import (
	"net/http"

	"util"

	"github.com/gin-gonic/gin"
)

func loginHandler(c *gin.Context) {
	payload := loginPayload{}
	c.BindJSON(&payload)

	user, err := FindUser(payload.Username)

	if err != nil {
		c.Error(util.CreateAPIError(http.StatusUnauthorized, err.Error()))
		return
	}

	err = ConfirmPassword(user, payload.Password)

	if err != nil {
		c.Error(util.CreateAPIError(http.StatusUnauthorized, "invalid credentials", err.Error()))
		return
	}

	jwt, err := generateJwt(user.ID)

	if err != nil {
		c.Error(util.CreateAPIError(http.StatusUnauthorized, "could not generate token", err.Error()))
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
		c.Error(util.CreateAPIError(http.StatusUnauthorized, "username has been take", "can not create user"))
		return
	}

	user := User{
		Username: payload.Username,
		Password: payload.Password,
	}

	err = CreateUser(&user)
	if err != nil {
		c.Error(util.CreateAPIError(http.StatusUnauthorized, err.Error(), "can not create user"))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user create successfully",
		"user":    user,
	})

}

func getUserHandler(c *gin.Context) {
	userID := c.GetString("UserID")
	user, err := FindUserByID(userID)
	if err != nil {
		c.Error(util.CreateAPIError(http.StatusUnauthorized, "could not find user", err.Error()))
		return
	}

	c.JSON(http.StatusOK, user)
}
