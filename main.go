package main

import (
	"log"

	database "github.com/sebito91/database/db"
)

func main() {
	db := database.NewDatabase()

	if err := db.Run(); err != nil {
		log.Fatal(err)
	}
}
