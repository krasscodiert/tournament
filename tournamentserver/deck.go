package tournamentserver

const (
	colorBlack     = "black"
	colorBlue      = "blue"
	colorWhite     = "white"
	colorRed       = "red"
	colorGreen     = "green"
	colorColorless = "colorless"
)

//Deck - Deck
type Deck struct {
	Name   string
	Colors map[string]int
	Cards  []string
}
