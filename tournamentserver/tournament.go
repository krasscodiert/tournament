package tournamentserver

//Tournament - Tournament struct
type Tournament struct {
	games        *graph
	participants []*User
}
