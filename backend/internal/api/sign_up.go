package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/cidstein/super-brunfo/internal/service"
)

func SignUp(db *pgx.Conn) gin.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
	}

	return func(c *gin.Context) {
		s := service.SignUpUseCase{}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(400, err.Error())
			return
		}

		if req.Email == "" {
			c.String(400, "email is required")
			return
		}

		if req.Password == "" {
			c.String(400, "password is required")
			return
		}

		if req.Nickname == "" {
			c.String(400, "nickname is required")
			return
		}

		err := s.SignUp(c.Request.Context(), db, req.Email, req.Password, req.Nickname)
		if err != nil {
			c.String(502, err.Error())
			return
		}

		c.String(200, "ok")
	}
}
