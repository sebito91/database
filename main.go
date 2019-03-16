package main

import (
	"log"

	"github.com/sebito91/database/db"
)

func main() {
	d := db.NewDatabase()

	if err := d.Run(); err != nil {
		log.Fatal(err)
	}
}
