package tournamentserver

type node struct {
	Connections []*node
	IsStart     bool
	IsEnd       bool
	Value       interface{}
}
