package api

import (
	"github.com/cidstein/super-brunfo/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func LoadRound(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := service.LoadRoundUseCase{}

		matchID := c.Query("match_id")
		if matchID == "" {
			c.String(400, "match_id is required")
			return
		}

		lr, err := s.LoadRound(c.Request.Context(), db, matchID)
		if err != nil {
			c.String(502, err.Error())
			return
		}

		if lr.ID == "" {
			c.String(404, "No round found")
			return
		}

		c.JSON(200, lr)
	}
}
