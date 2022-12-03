package entity

type MatchRepositoryInterface interface {
	Save(match Match) error
}
