package api

import "github.com/gin-gonic/gin"

func Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "Super Brunfo!")
	}
}
