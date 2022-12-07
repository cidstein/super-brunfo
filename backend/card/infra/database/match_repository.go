package database

import (
	"context"

	"github.com/cidstein/super-brunfo/card/entity"
	"github.com/jackc/pgx/v5"
)

type MatchRepository struct {
	Db *pgx.Conn
}

func NewMatchRepository(db *pgx.Conn) *MatchRepository {
	return &MatchRepository{Db: db}
}

func (r *MatchRepository) Save(ctx context.Context, match entity.Match) error {
	_, err := r.Db.Exec(
		ctx,
		"INSERT INTO matches (id, deck_player_id, deck_com_id, victory, finished) VALUES ($1, $2, $3, $4, $5)",
		match.ID,
		match.PlayerDeckID,
		match.NpcDeckID,
		match.Victory,
		match.Finished,
	)

	return err
}

func (r *MatchRepository) Update(ctx context.Context, match entity.Match) error {
	_, err := r.Db.Exec(
		ctx,
		"UPDATE matches SET deck_player_id = $1, deck_com_id = $2, victory = $3, finished = $4 WHERE id = $5",
		match.PlayerDeckID,
		match.NpcDeckID,
		match.Victory,
		match.Finished,
		match.ID,
	)

	return err
}

func (r *MatchRepository) Delete(ctx context.Context, id string) error {
	_, err := r.Db.Exec(
		ctx,
		"DELETE FROM matches WHERE id = $1",
		id,
	)

	return err
}

func (r *MatchRepository) FindByID(ctx context.Context, id string) (entity.Match, error) {
	var match entity.Match

	err := r.Db.QueryRow(
		ctx,
		"SELECT id, deck_player_id, deck_com_id, victory, finished FROM matches WHERE id = $1",
		id,
	).Scan(&match.ID, &match.PlayerDeckID, &match.NpcDeckID, &match.Victory, &match.Finished)

	return match, err
}

func (r *MatchRepository) ComputeWinner(ctx context.Context, match entity.Match) (entity.Match, error) {
	var rounds, roundWons int

	err := r.Db.QueryRow(
		ctx,
		"SELECT count(*) FROM round WHERE match_id = $1",
		match.ID,
	).Scan(&rounds)
	if err != nil {
		return match, err
	}

	err = r.Db.QueryRow(
		ctx,
		"SELECT count(*) FROM round WHERE match_id = $1 AND victory",
		match.ID,
	).Scan(&roundWons)
	if err != nil {
		return match, err
	}

	if roundWons > rounds/2 {
		match.Victory = true
	} else {
		match.Victory = false
	}

	match.Finished = true

	r.Update(ctx, match)

	return match, nil
}
