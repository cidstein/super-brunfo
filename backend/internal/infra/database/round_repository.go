package database

import (
	"context"

	"github.com/cidstein/super-brunfo/internal/model"
	"github.com/jackc/pgx/v5"
)

type RoundRepositoryInterface interface {
	Save(ctx context.Context, round model.Round) error
	Update(ctx context.Context, round model.Round) error
	FindCardsByID(ctx context.Context, id string) ([]model.Card, error)
}

type RoundRepository struct {
	Db *pgx.Conn
}

func NewRoundRepository(db *pgx.Conn) *RoundRepository {
	return &RoundRepository{Db: db}
}

func (r *RoundRepository) Save(ctx context.Context, round model.Round) error {
	_, err := r.Db.Exec(
		ctx,
		"INSERT INTO round (id, match_id, player_card_id, npc_card_id, victory, attribute) VALUES ($1, $2, $3, $4, $5, $6)",
		round.ID,
		round.MatchID,
		round.PlayerCardID,
		round.NpcCardID,
		round.Victory,
		round.Attribute,
	)

	return err
}

func (r *RoundRepository) Update(ctx context.Context, round model.Round) error {
	_, err := r.Db.Exec(
		ctx,
		"UPDATE round SET match_id = $1, player_card_id = $2, npc_card_id = $3, victory = $4, attribute = $5 WHERE id = $6",
		round.MatchID,
		round.PlayerCardID,
		round.NpcCardID,
		round.Victory,
		round.Attribute,
		round.ID,
	)

	return err
}

func (r *RoundRepository) FindCardsByID(ctx context.Context, id string) ([]model.Card, error) {
	rows, err := r.Db.Query(
		ctx,
		`
			select
				s.id,
				s.name,
				s.attack,
				s.defense,
				s.intelligence,
				s.agility,
				s.resilience,
				s.image_url
			from
			(
				select
					c.id,
					c.name,
					c.attack,
					c.defense,
					c.intelligence,
					c.agility,
					c.resilience,
					c.image_url,
					1 _order
				from
					round r
					join card c on
						r.player_card_id = c.id
				where
					r.id = $1
				union
				select
					c.id,
					c.name,
					c.attack,
					c.defense,
					c.intelligence,
					c.agility,
					c.resilience,
					c.image_url,
					2 _order
				from
					round r
					join card c on
						r.npc_card_id = c.id
				where
					r.id = $1
			) s
			order by
				s._order
		`,
		id,
	)
	if err != nil {
		return nil, err
	}

	var cards []model.Card
	for rows.Next() {
		var card model.Card
		err = rows.Scan(
			&card.ID,
			&card.Name,
			&card.Attack,
			&card.Defense,
			&card.Intelligence,
			&card.Agility,
			&card.Resilience,
			&card.ImageURL,
		)
		if err != nil {
			return nil, err
		}

		cards = append(cards, card)
	}

	return cards, nil
}
