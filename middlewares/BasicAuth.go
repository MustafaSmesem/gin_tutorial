package middlewares

import (
	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"mustafa": "m1234",
		"joud":    "j1234",
	})
}
