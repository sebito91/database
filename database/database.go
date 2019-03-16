package db

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// NewDatabase returns a new instance of our Database
func NewDatabase() *Database {
	return &Database{
		root: &DBValues{make(map[string]string)},
	}
}

// Run will actually kick off our database
func (d *Database) Run() error {
	return d.run()
}

func (d *Database) run() error {
	var bits []string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf(">> ")
	for scanner.Scan() {
		bits = strings.Split(scanner.Text(), " ")

		switch strings.ToUpper(bits[0]) {
		case "BEGIN":
			fmt.Printf("caught a begin: %s\n", bits)
		case "SET":
			fmt.Printf("caught a set: %s\n", bits)
			if len(bits) != 3 {
				fmt.Printf("incorrect number of arguments for set, expected 3, got %d: %s", len(bits), bits)
				break
			}
			d.Add(bits)
		case "GET":
			fmt.Printf("caught a get: %s\n", bits)
		case "DELETE":
			fmt.Printf("caught a delete: %s\n", bits)
		case "COUNT":
			fmt.Printf("caught a count: %s\n", bits)
		case "END":
			fmt.Printf("caught the end: %s\n", bits)
			return nil
		case "ROLLBACK":
			fmt.Printf("caught a rollback: %s\n", bits)
		case "COMMIT":
			fmt.Printf("caught a commit: %s\n", bits)
		case "":
		default:
			fmt.Printf("received unrecognized instruction: %s\n", strings.Join(bits, " "))
		}

		d.Print()
		fmt.Printf(">> ")
	}

	return fmt.Errorf("could not scan: %s", scanner.Err().Error())
}
