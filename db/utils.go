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
	if val, ok := d.vals[key]; ok {
		return val
	}

	t := d.next

	for {
		if t == nil {
			break
		}

		if val, ok := t.vals[key]; ok {
			return val
		}

		t = t.next
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

	var count int
	for {
		count++
		fmt.Printf("DEBUG -- %d. t: %+v, d: %+v\n", count, t, d)

		for k, v := range t.vals {
			fmt.Printf("DEBUG -- updating %s, %s\n", k, v)
			d.vals[k] = v
		}

		if t.prev == nil {
			fmt.Printf("breaking after %d runs...\n", count)
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
