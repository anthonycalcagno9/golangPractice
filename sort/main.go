package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

// Got this code from standard library sort package examples
// sort.Sort takes in an interface that implements Len, Swap, and Less methods
// We create a new type ByAge that implements these methods for sorting by age
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// Something I noticed when trying to sort by first name
// Swap and Less parameters are type int and I tried to change those to string
// but those are referenceing the indexes of the slice which are always int
// And name comes into play in the Less method
type ByFirst []Person

func (f ByFirst) Len() int           { return len(f) }
func (f ByFirst) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f ByFirst) Less(i, j int) bool { return f[i].First < f[j].First }

func main() {
	// Create a slice of strings
	fruits := []string{"banana", "apple", "cherry"}

	// Sort the slice
	sort.Strings(fruits)

	// Print the sorted slice
	fmt.Println(fruits)

	p1 := Person{
		First: "anthony",
		Age:   29,
	}
	p2 := Person{
		First: "rafa",
		Age:   28,
	}
	p3 := Person{
		First: "bibs",
		Age:   27,
	}
	p4 := Person{
		First: "mike",
		Age:   30,
	}

	people := []Person{p1, p2, p3, p4}

	fmt.Printf("people = %v\n", people)

	sort.Sort(ByAge(people))
	fmt.Printf("people sorted by age = %v\n", people)

	sort.Sort(ByFirst(people))
	fmt.Printf("people sorted by first name = %v\n", people)

}
