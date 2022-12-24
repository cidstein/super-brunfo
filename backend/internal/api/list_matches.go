package api

import (
	"github.com/cidstein/super-brunfo/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func ListMatches(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := service.ListMatchesUseCase{}
		lm, err := s.ListMatches(c.Request.Context(), db)
		if err != nil {
			c.String(500, err.Error())
			return
		}

		if len(lm) == 0 {
			c.String(404, "No matches found")
			return
		}

		c.JSON(200, lm)
	}
}
