package main

import (
	"log"

	"github.com/sebito91/database/db"
)

var root *db.Database

func init() {
	root = db.NewDatabase()
}

func main() {
	if err := root.Run(); err != nil {
		log.Fatal(err)
	}
}
