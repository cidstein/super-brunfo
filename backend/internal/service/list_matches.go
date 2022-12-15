package service

import (
	"context"

	"github.com/cidstein/super-brunfo/internal/infra/database"
)

type ListMatchesUseCase struct {
	MatchRepository database.MatchRepositoryInterface
}

func (s *ListMatchesUseCase) ListMatches(ctx context.Context) ([]MatchOutputDTO, error) {
	matches, err := s.MatchRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	matchesDTO := make([]MatchOutputDTO, len(matches))
	for i, match := range matches {
		matchesDTO[i] = MatchOutputDTO{
			ID:       match.ID,
			Victory:  match.Victory,
			Finished: match.Finished,
		}
	}

	return matchesDTO, nil
}
