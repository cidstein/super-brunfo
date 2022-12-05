package database

import (
	"database/sql"

	"github.com/cidstein/super-brunfo/card/entity"
)

type RoundRepository struct {
	Db *sql.DB
}

func NewRoundRepository(db *sql.DB) *RoundRepository {
	return &RoundRepository{Db: db}
}

func (r *RoundRepository) Save(round entity.Round) error {
	_, err := r.Db.Exec(
		"INSERT INTO rounds (id, match_id, player_card_id, npc_card_id, victory, attribute) VALUES ($1, $2, $3, $4, $5, $6)",
		round.ID,
		round.MatchID,
		round.PlayerCardID,
		round.NpcCardID,
		round.Victory,
		round.Attribute,
	)

	return err
}

func (r *RoundRepository) Update(round entity.Round) error {
	_, err := r.Db.Exec(
		"UPDATE rounds SET match_id = $1, player_card_id = $2, npc_card_id = $3, victory = $4, attribute = $5 WHERE id = $6",
		round.MatchID,
		round.PlayerCardID,
		round.NpcCardID,
		round.Victory,
		round.Attribute,
		round.ID,
	)

	return err
}
