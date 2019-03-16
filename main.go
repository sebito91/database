package main

import (
	"log"
)

func main() {
	db := database.NewDatabase()

	if err := db.Run(); err != nil {
		log.Fatal(err)
	}
}
