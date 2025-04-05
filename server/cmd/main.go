package main

import (
	"log"

	"github.com/dipenkumarr/go-next-chat/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Could not initialize database: ", err)
	}
}
