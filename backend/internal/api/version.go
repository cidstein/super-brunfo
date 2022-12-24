package api

import "github.com/gin-gonic/gin"

func Version(version string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, version)
	}
}
