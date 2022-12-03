package entity

type CardRepositotyInterface interface {
	FindAll() ([]Card, error)
}

type DeckRepositoryInterface interface {
	FindAll() ([]Deck, error)
}

type MatchRepositoryInterface interface {
	Save(match Match) error
}
