package database

import (
	"database/sql"

	"github.com/cidstein/super-brunfo/card/entity"
)

type MatchRepository struct {
	Db *sql.DB
}

func NewMatchRepository(db *sql.DB) *MatchRepository {
	return &MatchRepository{Db: db}
}

func (r *MatchRepository) Save(match entity.Match) error {
	_, err := r.Db.Exec(
		"INSERT INTO matches (id, deck_player_id, deck_com_id, victory) VALUES ($1, $2, $3, $4)",
		match.ID,
		match.PlayerDeckID,
		match.NpcDeckID,
		match.Victory,
	)

	return err
}

func (r *MatchRepository) Update(match entity.Match) error {
	_, err := r.Db.Exec(
		"UPDATE matches SET deck_player_id = $1, deck_com_id = $2, victory = $3 WHERE id = $4",
		match.PlayerDeckID,
		match.NpcDeckID,
		match.Victory,
		match.ID,
	)

	return err
}

func (r *MatchRepository) FindByID(id string) (entity.Match, error) {
	var match entity.Match

	err := r.Db.QueryRow(
		"SELECT id, deck_player_id, deck_com_id, victory FROM matches WHERE id = $1",
		id,
	).Scan(&match.ID, &match.PlayerDeckID, &match.NpcDeckID, &match.Victory)

	return match, err
}

func (r *MatchRepository) ComputeWinner(match entity.Match) (entity.Match, error) {
	var rounds, roundWons int

	err := r.Db.QueryRow(
		"SELECT count(*) FROM round WHERE match_id = $1",
		match.ID,
	).Scan(&rounds)
	if err != nil {
		return match, err
	}

	err = r.Db.QueryRow(
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

	r.Update(match)

	return match, nil
}
