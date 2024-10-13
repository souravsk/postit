package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/souravsk/go-zero-to-hero/user_auth/utils"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(cookie)

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("User", claims.User)
		c.Next()
	}
}
