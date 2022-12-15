package api

import (
	"encoding/json"
	"net/http"

	"github.com/cidstein/super-brunfo/internal/service"
	"github.com/jackc/pgx/v5"
)

type request struct {
	MatchID   string `json:"match_id" validate:"required"`
	Attribute string `json:"attribute" validate:"required"`
}

func PlayGame(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Method = http.MethodPost

		pguc := service.PlayGameUseCase{}

		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		matchID := req.MatchID
		if matchID == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("match_id is required"))
			return
		}

		attribute := req.Attribute
		if attribute == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("attribute is required"))
			return
		}

		pg, err := pguc.Play(r.Context(), db, matchID, attribute)
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
