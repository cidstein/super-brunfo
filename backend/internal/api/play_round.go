package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cidstein/super-brunfo/internal/service"
	"github.com/jackc/pgx/v5"
)

func PlayRound(db *pgx.Conn) http.HandlerFunc {
	type request struct {
		RoundID   string `json:"round_id"`
		Attribute string `json:"attribute"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		log.Printf("PlayRound")
		log.Print("preflight detected", r.Header)

		pruc := service.PlayRoundUseCase{}

		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		roundID := req.RoundID
		if roundID == "" {
			log.Printf("round_id is required")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("round_id is required"))
			return
		}

		attribute := req.Attribute
		if attribute == "" {
			log.Printf("attribute is required")
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
