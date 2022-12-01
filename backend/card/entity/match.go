package entity

type Match struct {
	ID     string
	Deck1  Deck
	Deck2  Deck
	Winner string
}

func NewMatch(id string, deck1 Deck, deck2 Deck) Match {
	return Match{
		ID:     id,
		Deck1:  deck1,
		Deck2:  deck2,
		Winner: "",
	}
}
