package tournamentserver

import (
	"github.com/krasscodiert/tournament/tournamentserver/idfactories"
)

// User - User struct
type User struct {
	ID           int
	Name         string
	Decks        []*Deck
	MatchHistory []string
}

// New - Default constructor
func (u *User) New(name string, idFactory idfactories.IDFactory) *User {
	p := new(User)
	p.Name = name
	p.ID = idFactory.GetID()
	p.MatchHistory = make([]string, 20)
	p.Decks = []*Deck{}
	return p
}
