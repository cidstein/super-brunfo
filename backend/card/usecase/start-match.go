package usecase

import "github.com/cidstein/super-brunfo/card/entity"

type CardOutputDTO struct {
	ID           string
	Name         string
	Attack       int
	Defense      int
	Intelligence int
	Agility      int
	Resilience   int
}

type DeckOutputDTO struct {
	ID    string
	Cards []CardOutputDTO
}

type MatchOutputDTO struct {
	ID     string
	Deck1  DeckOutputDTO
	Deck2  DeckOutputDTO
	Winner string
}

type StartMatchUseCase struct {
	MatchRepository entity.MatchRepositoryInterface
}

// func (s *StartMatchUseCase) Start(deck1 entity.Deck, deck2 entity.Deck) (MatchOutputDTO, error) {
// 	cards := s.MatchRepository.FindAll()
// 	if len(cards) < 2 {
// 		return MatchOutputDTO{}, nil
// 	}

// 	match := entity.NewMatch(deck1, deck2)
// 	match.Start()

// 	err := s.MatchRepository.Save(match)
// 	if err != nil {
// 		return MatchOutputDTO{}, err
// 	}

// 	return MatchOutputDTO{
// 		ID:     match.ID,
// 		Deck1:  convertDeckToOutputDTO(match.Deck1),
// 		Deck2:  convertDeckToOutputDTO(match.Deck2),
// 		Winner: match.Winner,
// 	}, nil
// }

func StartMatch() (*MatchOutputDTO, error) {
	/*
		Embaralhar deck
		Atribuir deck para cada jogador
		Definir quem começa
	*/

	return nil, nil
}
