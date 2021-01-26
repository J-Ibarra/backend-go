module main

go 1.15

replace (
	auth => ./auth
	util => ./util
)

require (
	auth v0.0.0
	github.com/gin-gonic/gin v1.6.3
)
