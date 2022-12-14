package service

type CardOutputDTO struct {
	ID           string
	Name         string
	Attack       int
	Defense      int
	Intelligence int
	Agility      int
	Resilience   int
	FlavourText  string
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
	Counter    int
	Victory    bool
	Finished   bool
}

type RoundOutputDTO struct {
	ID        string
	Match     MatchOutputDTO
	Cards     []CardOutputDTO
	Counter   int
	Victory   bool
	Attribute string
}
