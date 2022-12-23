package database

import (
	"context"

	"github.com/cidstein/super-brunfo/internal/model"
	"github.com/jackc/pgx/v5"
)

type RoundRepositoryInterface interface {
	Save(ctx context.Context, round model.Round) error
	Update(ctx context.Context, round model.Round) error
	FindByID(ctx context.Context, id string) (*model.Round, error)
	FindRoundToBePlayed(ctx context.Context, matchID string) (*model.Round, error)
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
		"INSERT INTO round (id, match_id, player_card_id, npc_card_id, counter, victory, attribute) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		round.ID,
		round.MatchID,
		round.PlayerCardID,
		round.NpcCardID,
		round.Counter,
		round.Victory,
		round.Attribute,
	)

	return err
}

func (r *RoundRepository) Update(ctx context.Context, round model.Round) error {
	_, err := r.Db.Exec(
		ctx,
		`
			UPDATE
				round
			SET
				match_id = $1,
				player_card_id = $2,
				npc_card_id = $3,
				counter = $4,
				victory = $5,
				attribute = $6
			WHERE
				id = $7
		`,
		round.MatchID,
		round.PlayerCardID,
		round.NpcCardID,
		round.Counter,
		round.Victory,
		round.Attribute,
		round.ID,
	)

	return err
}

func (r *RoundRepository) FindByID(ctx context.Context, id string) (*model.Round, error) {
	round := model.Round{}
	err := r.Db.QueryRow(
		ctx,
		"SELECT id, match_id, player_card_id, npc_card_id, counter, victory, attribute FROM round WHERE id = $1",
		id,
	).Scan(
		&round.ID,
		&round.MatchID,
		&round.PlayerCardID,
		&round.NpcCardID,
		&round.Counter,
		&round.Victory,
		&round.Attribute,
	)

	if err != nil {
		return nil, err
	}

	return &round, nil
}

func (r *RoundRepository) FindRoundToBePlayed(ctx context.Context, matchID string) (*model.Round, error) {
	round := model.Round{}
	err := r.Db.QueryRow(
		ctx,
		`
			select
				r.id,
				r.match_id,
				r.player_card_id,
				r.npc_card_id,
				r.counter,
				r.victory,
				r.attribute
			from
				round r
				join match m on
					r.match_id = m.id
			where
				r.match_id = $1
				and nullif(r.attribute, '') is null
			order by
				r.counter
			limit 1
		`,
		matchID,
	).Scan(
		&round.ID,
		&round.MatchID,
		&round.PlayerCardID,
		&round.NpcCardID,
		&round.Counter,
		&round.Victory,
		&round.Attribute,
	)

	if err != nil {
		return nil, err
	}

	return &round, nil
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
