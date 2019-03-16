package db

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// NewDatabase returns a new instance of our Database
func NewDatabase() *Database {
	return &Database{}
}

// Run will actually kick off our database
func (d *Database) Run() error {
	return d.run()
}

// errorMsg is just a func for us to reduce repetition
func errorMsg(bits []string, exp int) {
	fmt.Printf("incorrect number of arguments for %s, expected %d, got %d: %s\n", strings.ToUpper(bits[0]), exp, len(bits), strings.Join(bits, " "))
}

func (d *Database) run() error {
	var bits []string

	scanner := bufio.NewScanner(os.Stdin)

	d.vals = make(map[string]string)

	fmt.Printf(">> ")
	for scanner.Scan() {
		bits = strings.Split(scanner.Text(), " ")

		switch strings.ToUpper(bits[0]) {
		case "BEGIN":
			t := d.next
			d.next = &Database{vals: make(map[string]string), next: t, prev: d}

			if t != nil {
				t.prev = d.next
			}
		case "SET":
			if len(bits) != 3 {
				errorMsg(bits, 3)
				break
			}

			d.Set(bits)
		case "GET":
			if len(bits) != 2 {
				errorMsg(bits, 2)
				break
			}

			fmt.Printf("%s\n", d.Get(bits[1]))
		case "DELETE":
			if len(bits) != 2 {
				errorMsg(bits, 2)
				break
			}

			d.Delete(bits[1])
		case "COUNT":
			if len(bits) != 2 {
				errorMsg(bits, 2)
				break
			}

			fmt.Printf("%d\n", d.Count(bits[1]))
		case "END":
			return nil
		case "ROLLBACK":
			if d.next == nil {
				fmt.Printf("TRANSACTION NOT FOUND\n")
				break
			}
			d.next = d.next.next
		case "COMMIT":
			d.Commit()
		case "":
		default:
			fmt.Printf("received unrecognized instruction: %s\n", strings.Join(bits, " "))
		}

		//		d.Print()
		fmt.Printf(">> ")
	}

	return fmt.Errorf("could not scan: %s", scanner.Err().Error())
}
