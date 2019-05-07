package tournament_server

type Tournament struct {
	games *graph
	participants []*Participant
}
