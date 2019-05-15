package tournamentserver

type graph struct {
	Beginner []*node
	Currents map[*node]int
}
