package main

import "fmt"

type youngin interface {
	Walk()
	Run()
}

type Dog struct {
	Name string
}

// value type receiver
func (d Dog) Walk() {
	fmt.Printf("%s is walking\n", d.Name)
}

// pointer type receiver
func (d *Dog) Run() {
	//Don't need to dereference d, Go does it automatically
	//(*d).Name = "Rover" is how you explicitly dereference
	d.Name = "Rover"
	fmt.Printf("%s is running\n", d.Name)
}

func younginRun(y youngin) {
	y.Run()
}

// In production code, you want to keep your methods consistent
// Either use value receivers or pointer receivers for each method on a certain type
func main() {
	dog := Dog{Name: "Buddy"}
	dog.Walk()
	dog.Run()

	// This will not work because dog is a value type and Run() expects a pointer type receiver
	//younginRun(dog)

	dog2 := &Dog{Name: "Max"}
	dog2.Walk()
	dog2.Run()
	// This will work even though Walk() is a value receiver and dog2 is a pointer type
	// because Go automatically dereferences the pointer when calling a method with a value receiver, so pointers are more flexible in this regard
	younginRun(dog2)
}
