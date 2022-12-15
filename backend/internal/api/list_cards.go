package api

import (
	"encoding/json"
	"net/http"

	"github.com/cidstein/super-brunfo/internal/service"
	"github.com/jackc/pgx/v5"
)

func ListCards(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Method = http.MethodGet

		lcuc := service.ListCardsUseCase{}

		lc, err := lcuc.ListCards(r.Context(), db)
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
