package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/cidstein/super-brunfo/internal/service"
)

func SignIn(db *pgx.Conn) gin.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(c *gin.Context) {
		s := service.SignInUseCase{}

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

		err := s.SignIn(c.Request.Context(), db, req.Email, req.Password)
		if err != nil {
			c.String(502, err.Error())
			return
		}

		ss := sessions.Default(c)
		ss.Set("email", req.Email)
		ss.Save()

		c.String(200, "User signed in successfully")
	}
}
