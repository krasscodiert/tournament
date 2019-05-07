package tournament_server

type node struct {
	Connections []*node
	IsStart bool
	IsEnd bool
	Value interface{}
}