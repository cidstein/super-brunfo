package entity

import "context"

type CardRepositoryInterface interface {
	Save(ctx context.Context, card Card) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (Card, error)
	FindAll(ctx context.Context) ([]Card, error)
}

type DeckRepositoryInterface interface {
	Save(ctx context.Context) (*Deck, error)
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (Deck, error)
}

type MatchRepositoryInterface interface {
	Save(ctx context.Context, match Match) error
	Update(ctx context.Context, match Match) error
	FindByID(ctx context.Context, id string) (Match, error)
	ComputeWinner(ctx context.Context, match Match) (Match, error)
}

type RoundRepositoryInterface interface {
	Save(ctx context.Context, round Round) error
	Update(ctx context.Context, round Round) error
}
