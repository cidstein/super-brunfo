package entity

type CardRepositotyInterface interface {
	FindAll() ([]Card, error)
}

type DeckRepositoryInterface interface {
	Save() (*Deck, error)
	FindByID(id string) (Deck, error)
}

type MatchRepositoryInterface interface {
	Save(match Match) error
	Update(match Match) error
	FindByID(id string) (Match, error)
	ComputeWinner(match Match) (Match, error)
}

type RoundRepositoryInterface interface {
	Save(round Round) error
	Update(round Round) error
}
