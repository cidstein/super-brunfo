package database

import (
	"context"

	"github.com/cidstein/super-brunfo/card/entity"
	"github.com/jackc/pgx/v5"
)

type RoundRepository struct {
	Db *pgx.Conn
}

func NewRoundRepository(db *pgx.Conn) *RoundRepository {
	return &RoundRepository{Db: db}
}

func (r *RoundRepository) Save(ctx context.Context, round entity.Round) error {
	_, err := r.Db.Exec(
		ctx,
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

func (r *RoundRepository) Update(ctx context.Context, round entity.Round) error {
	_, err := r.Db.Exec(
		ctx,
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
