package tournament_server

const(
	COLOR_BLACK = "black"
	COLOR_BLUE = "blue"
	COLOR_WHITE = "white"
	COLOR_RED = "red"
	COLOR_GREEN = "green"
	COLOR_COLORLESS = "colorless"
)

type Deck struct {
	Name string
	Colors map[string]int
	Cards []string
}