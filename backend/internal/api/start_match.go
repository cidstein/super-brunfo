package api

import (
	"github.com/cidstein/super-brunfo/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func StartMatch(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := service.StartMatchUseCase{}

		sm, err := s.Start(c.Request.Context(), db)
		if err != nil {
			c.String(502, err.Error())
			return
		}

		c.JSON(200, sm)
	}
}
