package main

import (
	"fmt"
	"log"
	"strconv"
)

type person struct {
	first string
}

type secretAgent struct {
	person person
	ltk    bool
}

type human interface {
	speak()
}

type book struct {
	title string
}

type count int

func main() {
	p := person{
		first: "anthony",
	}

	sa := secretAgent{
		person: p,
		ltk:    true,
	}

	p.speak()
	sa.speak()
	sa.person.speak()

	//polymorphism is implemented here. Both person and secretAgent can be passed in to this function
	//because they are both of type human, since they both have the speak() method
	saySomething(p)
	saySomething(sa)

	//Working with Stringer interface
	//type Stringer interface {
	//    String() string
	//}
	b := book{
		title: "game of thrones",
	}

	var n count = 42

	log.Println(b)
	log.Println(n)

	logInfo(b)
	logInfo(n)

	fmt.Println("------------------------------------------")

}

func (p person) speak() {
	fmt.Println("I am a person", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.person.first)
}

func saySomething(h human) {
	h.speak()
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

func (c count) String() string {
	return fmt.Sprint("this is the number ", strconv.Itoa(int(c)))
}

// polymorphism since anything that satisfies the Stringer interface can be used here
func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM 138", s.String())
}
