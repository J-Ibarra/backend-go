package main

import (
	"./auth"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	auth.ConfigureRoutes(r)

	r.Run()
}
