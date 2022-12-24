package api

import (
	"github.com/cidstein/super-brunfo/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func PlayRound(db *pgx.Conn) gin.HandlerFunc {
	type request struct {
		RoundID   string `json:"round_id"`
		Attribute string `json:"attribute"`
	}

	return func(c *gin.Context) {
		s := service.PlayRoundUseCase{}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(400, err.Error())
			return
		}

		if req.RoundID == "" {
			c.String(400, "round_id is required")
			return
		}

		if req.Attribute == "" {
			c.String(400, "attribute is required")
			return
		}

		pg, err := s.Play(c.Request.Context(), db, req.RoundID, req.Attribute)
		if err != nil {
			c.String(502, err.Error())
			return
		}

		c.JSON(200, pg)
	}
}
