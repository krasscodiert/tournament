package tournament_server

type graph struct {
	Beginner []*node
	Currents map[*node]int
}
