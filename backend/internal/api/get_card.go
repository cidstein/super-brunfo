package api

import (
	"encoding/json"
	"net/http"

	"github.com/cidstein/super-brunfo/internal/service"
	"github.com/jackc/pgx/v5"
)

func GetCard(db *pgx.Conn) http.HandlerFunc {
	type request struct {
		CardID string `json:"card_id" validate:"required"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		r.Method = http.MethodGet

		gcuc := service.GetCardUseCase{}

		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		cardID := req.CardID
		if cardID == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("card_id is required"))
			return
		}

		lc, err := gcuc.GetCard(r.Context(), db, cardID)
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
