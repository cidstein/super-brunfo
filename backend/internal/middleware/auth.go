package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ss := sessions.Default(c)
		ssID := ss.Get("email")

		if ssID == nil {
			c.String(401, "Unauthorized")
			c.Abort()
			return
		}
	}
}
