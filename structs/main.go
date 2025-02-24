package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type husband struct {
	person       person
	lostRing     bool
	favoriteGame string
}

type wife struct {
	person         person
	lostRing       bool
	favoriteFlower string
}

func main() {
	p1 := person{
		first: "anthony",
		last:  "calcagno",
		age:   29,
	}

	p2 := person{
		first: "rafa",
		last:  "santos",
		age:   28,
	}

	fmt.Printf("printing out one value of struct = %v\n", p1.first)

	fmt.Printf("printing out whole struct = %v\n", p2)

	fmt.Printf("Type = %T and value = %v\n", p1, p1)

	//Embedded structs
	//Inner fields in person get PROMOTED to both husband and wife
	hus := husband{
		person: person{
			first: "ant",
			last:  "cal",
			age:   29,
		},
		lostRing:     true,
		favoriteGame: "fire emblem",
	}

	wife := wife{
		person: person{
			first: "rafa",
			last:  "sant",
			age:   28,
		},
		lostRing:       false,
		favoriteFlower: "rose",
	}

	fmt.Println(wife)
	fmt.Println(wife.person.first)
	fmt.Println(hus.lostRing)
	fmt.Println(hus.favoriteGame)

	//Anonymous struct
	//Can be useful when you don't need a reusable struct, just need it in one place
	jot := struct {
		first string
		last  string
		age   int
	}{
		first: "jot",
		last:  "buj",
		age:   29,
	}

	fmt.Printf("printing out anonymous struct type = %T\n", jot)

	//you can create a type in go

	type foo string

	//you might do this ^^^ because you can add methods to foo
	//so now a basic string type has extra methods that you can control
	//the time package does this with int64

	//COOL FACT - For best performance, you want to organize your struct from largest datatype to smallest from top to bottom

}
