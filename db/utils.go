package db

import "fmt"

// Database is the stack implementation of our set of databases
type Database struct {
	prev *Database
	next *Database

	vals map[string]string
}

// Set pushes the elements to the values map
func (d *Database) Set(bits []string) {
	t := d
	if d.next != nil {
		t = d.next
	}

	t.vals[bits[1]] = bits[2]
}

// Get the value for the provided key
func (d *Database) Get(key string) string {
	t := d
	for {
		if t.next == nil {
			break
		}

		t = t.next
		if val, ok := t.vals[key]; ok {
			return val
		}
	}

	if val, ok := d.vals[key]; ok {
		return val
	}

	return "NULL"
}

// Count returns the number of instances of the given value
func (d *Database) Count(val string) int {
	t := d
	if t.next != nil {
		t = t.next
	}

	var count int
	for _, v := range t.vals {
		if v == val {
			count++
		}
	}
	return count
}

// Delete removes the provided key fromo the database
func (d *Database) Delete(key string) {
	t := d
	if d.next != nil {
		t = d.next
	}

	if _, ok := t.vals[key]; ok {
		delete(t.vals, key)
	}
}

// Commit the open transactions
func (d *Database) Commit() {
	if d.next == nil {
		return
	}

	t := d.next
	for {
		if t.next == nil {
			break
		}
		t = t.next
	}

	for {
		for k, v := range t.vals {
			d.vals[k] = v
		}

		if t.prev == nil {
			break
		}

		t = t.prev
	}

	d.next = nil
}

// Print out the elements of the database
func (d *Database) Print() {
	for k, v := range d.vals {
		fmt.Printf("root k: %s, v: %s\n", k, v)
	}
}
