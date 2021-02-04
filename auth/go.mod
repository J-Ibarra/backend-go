module auth

go 1.15

replace util => ../util

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	gorm.io/gorm v1.20.12
	util v0.0.0
)
