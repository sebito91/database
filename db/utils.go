package db

import "fmt"

// Database is the stack implementation of our set of databases
type Database struct {
	root *Database
	next *Database

	vals map[string]string
}

// Set pushes the elements to the values map
func (d *Database) Set(bits []string) {
	d.vals[bits[1]] = bits[2]
}

// Get the value for the provided key
func (d *Database) Get(key string) string {
	if val, ok := d.vals[key]; ok {
		return val
	}
	return "NULL"
}

// Count returns the number of instances of the given value
func (d *Database) Count(val string) int {
	var count int
	for _, v := range d.vals {
		if v == val {
			count++
		}
	}
	return count
}

// Delete removes the provided key fromo the database
func (d *Database) Delete(key string) {
	if _, ok := d.vals[key]; ok {
		delete(d.vals, key)
	}
}

// Print out the elements of the database
func (d *Database) Print() {
	for k, v := range d.vals {
		fmt.Printf("root k: %s, v: %s\n", k, v)
	}
}
