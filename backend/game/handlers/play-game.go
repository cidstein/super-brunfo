package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cidstein/super-brunfo/game/usecases"
	"github.com/jackc/pgx/v5"
)

type request struct {
	MatchID   string `json:"match_id" validate:"required"`
	Attribute string `json:"attribute" validate:"required"`
}

func PlayGame(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pguc := usecases.PlayGameUseCase{}

		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		matchID := req.MatchID
		attribute := req.Attribute

		pg, err := pguc.Play(r.Context(), db, matchID, attribute)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		res, err := json.Marshal(pg)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(res)
	}
}
