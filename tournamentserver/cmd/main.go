package main

import (
	"fmt"

	"github.com/krasscodiert/tournament/tournamentserver"
)

func main() {
	db := new(tournamentserver.DB).New()
	api := new(tournamentserver.API).New(db)
	api.Start()
	defer db.Close()
	fmt.Println("hallo")
}
