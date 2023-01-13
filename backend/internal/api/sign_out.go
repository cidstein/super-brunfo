package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SignOut(db *pgx.Conn) gin.HandlerFunc {
	type request struct {
		Email string `json:"email"`
	}

	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(400, err.Error())
			return
		}

		if req.Email == "" {
			c.String(400, "email is required")
			return
		}

		ss := sessions.Default(c)
		ss.Clear()
		ss.Save()

		c.String(200, "User signed out successfully")
	}
}
