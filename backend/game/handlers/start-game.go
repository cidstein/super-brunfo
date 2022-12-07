package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cidstein/super-brunfo/game/usecases"
	"github.com/jackc/pgx/v5"
)

func StartMatch(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		smuc := usecases.StartMatchUseCase{}

		sm, err := smuc.Start(r.Context(), db)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		res, err := json.Marshal(sm)
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
