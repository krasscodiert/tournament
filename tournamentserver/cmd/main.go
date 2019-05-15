package main

import (
	"fmt"

	"github.com/krasscodiert/tournament/tournamentserver"
)

func main() {
	api := new(tournamentserver.API).New(new(tournamentserver.DB).New())
	api.Start()
	fmt.Println("hallo")
}
