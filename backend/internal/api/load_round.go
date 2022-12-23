package api

import (
	"encoding/json"
	"net/http"

	"github.com/cidstein/super-brunfo/internal/service"
	"github.com/jackc/pgx/v5"
)

func LoadRound(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		r.Method = http.MethodGet

		lruc := service.LoadRoundUseCase{}

		matchID := r.URL.Query().Get("match_id")
		if matchID == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("match_id is required"))
			return
		}

		lr, err := lruc.LoadRound(r.Context(), db, matchID)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte(err.Error()))
			return
		}

		res, err := json.Marshal(lr)
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
