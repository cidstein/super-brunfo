package entity

type Round struct {
	ID        string
	Match     Match
	Card1     Card
	Card2     Card
	Winner    string
	Atributte string
}
