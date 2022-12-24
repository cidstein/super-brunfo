package api

import (
	"github.com/cidstein/super-brunfo/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func ListCards(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := service.ListCardsUseCase{}
		lc, err := s.ListCards(c.Request.Context(), db)
		if err != nil {
			c.String(500, err.Error())
			return
		}

		if len(lc) == 0 {
			c.String(404, "No cards found")
			return
		}

		c.JSON(200, lc)
	}
}
