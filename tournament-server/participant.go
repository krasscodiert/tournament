package tournament_server

import "tournament-server/IDFactories"

type Participant struct {
	ID int
	Name string
	Decks []*Deck
	MatchHistory []string
}

func CreateParticipant(name string, idFactory IDFactories.IDFactory) *Participant {
	p := new(Participant)
	p.Name = name
	p.ID = idFactory.GetID()
	p.MatchHistory = make([]string, 20)
	p.Decks = make([]*Deck, 20)
	return p
}