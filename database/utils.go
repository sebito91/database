package db

import "fmt"

// DBValues is the default instance of the table, containing string keys and string values only
type DBValues struct {
	vals map[string]string
}

// Database is the stack implementation of our set of databases
type Database struct {
	root *DBValues
	next *DBValues
}

// Add pushes the elements to the values map
func (d *Database) Add(bits []string) {
	d.root.vals[bits[1]] = bits[2]
}

// Print out the elements of the database
func (d *Database) Print() {
	for k, v := range d.root.vals {
		fmt.Printf("root k: %s, v: %s\n", k, v)
	}
}
