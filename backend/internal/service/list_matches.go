package service

import (
	"context"

	"github.com/cidstein/super-brunfo/internal/infra/database"
	"github.com/jackc/pgx/v5"
)

type ListMatchesUseCase struct {
	MatchRepository database.MatchRepositoryInterface
}

func (s *ListMatchesUseCase) ListMatches(ctx context.Context, db *pgx.Conn) ([]MatchOutputDTO, error) {
	s.MatchRepository = database.NewMatchRepository(db)
	matches, err := s.MatchRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	matchesDTO := make([]MatchOutputDTO, len(matches))
	for i, match := range matches {
		matchesDTO[i] = MatchOutputDTO{
			ID:       match.ID,
			Counter:  match.Counter,
			Victory:  match.Victory,
			Finished: match.Finished,
		}
	}

	return matchesDTO, nil
}
