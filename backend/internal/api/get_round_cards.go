package api

import (
	"encoding/json"
	"net/http"

	"github.com/cidstein/super-brunfo/internal/service"
	"github.com/jackc/pgx/v5"
)

func GetRoundCards(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		r.Method = http.MethodGet

		lcuc := service.GetRoundCardsUseCase{}

		roundID := r.URL.Query().Get("id")
		if roundID == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id is required"))
			return
		}

		lc, err := lcuc.GetRoundCards(r.Context(), db, roundID)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte(err.Error()))
			return
		}

		res, err := json.Marshal(lc)
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
