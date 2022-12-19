package service

type CardOutputDTO struct {
	ID           string
	Name         string
	Attack       int
	Defense      int
	Intelligence int
	Agility      int
	Resilience   int
	ImageURL     string
}

type DeckOutputDTO struct {
	ID    string
	Cards []CardOutputDTO
}

type MatchOutputDTO struct {
	ID         string
	PlayerDeck DeckOutputDTO
	NpcDeck    DeckOutputDTO
	Victory    bool
	Finished   bool
}

type RoundOutputDTO struct {
	ID         string
	Match      MatchOutputDTO
	PlayerCard CardOutputDTO
	NpcCard    CardOutputDTO
	Victory    bool
	Attribute  string
}
