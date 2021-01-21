package main

import (
	"auth"

	"github.com/gin-gonic/gin"
)

func runServer(router *gin.Engine) {
	if mode := gin.Mode(); mode != gin.ReleaseMode {
		router.Run("127.0.0.1:8080")
	} else {
		router.Run()
	}
}

func main() {
	r := gin.Default()
	defer runServer(r)

	auth.ConfigureRoutes(r)
}
