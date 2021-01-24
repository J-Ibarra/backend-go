module src

go 1.15

require (
	auth v0.0.0
	github.com/gin-gonic/gin v1.6.3
)

replace auth => ./src/auth
