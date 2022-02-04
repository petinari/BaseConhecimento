package middlewares

import (
	"api/auth"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.FullPath() != "/login" {
			erro := auth.ValidarToken(c.Request)
			if erro != nil {
				c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
			}
			c.Next()
		}
	}
}
