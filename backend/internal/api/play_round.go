package api

import (
	"encoding/json"
	"net/http"

	"github.com/cidstein/super-brunfo/internal/service"
	"github.com/jackc/pgx/v5"
)

func PlayRound(db *pgx.Conn) http.HandlerFunc {
	type request struct {
		RoundID   string `json:"round_id" validate:"required"`
		Attribute string `json:"attribute" validate:"required"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		r.Method = http.MethodPut

		pruc := service.PlayRoundUseCase{}

		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		roundID := req.RoundID
		if roundID == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("round_id is required"))
			return
		}

		attribute := req.Attribute
		if attribute == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("attribute is required"))
			return
		}

		pg, err := pruc.Play(r.Context(), db, roundID, attribute)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte(err.Error()))
			return
		}

		res, err := json.Marshal(pg)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
