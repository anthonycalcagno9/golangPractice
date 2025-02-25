package main

import "fmt"

//named function
//func (r receiver) identifier(p parameters) (r returns) {code}

//anonymous function
//func(p parameters) (returns) {code} (pass in the actual parameters here, first parameter section is function signature)

//anonymous function can help to solve small task and limit scope in the middle of the code

func anonymousFunctionPractice() {

	func() {
		fmt.Println("hello world")
	}()

	func(s string) {
		fmt.Println(s)
	}("passing in parameter to anonymous function")

	s := func(s string) string {
		return s
	}("passing in parameter to anonymous function that also has return value")

	fmt.Println(s)

	//functions are first class citizens in go, so they act like any other normal type (ie, boolean, numeric, slice, structs, etc)
	//first class citizens can be assigned to variables, passed in to funcs as parameters, and be returned from functions
	//You can assign a function to a variable
	x := func() {
		fmt.Println("printing out function that was set to variable x")
	}

	y := func(s string) {
		fmt.Println("printing out func that was set to y, y = ", s)
	}

	x()
	y("passing into y")

	//Returning a func
	z := foop()
	fmt.Println(z())

}

func foop() func() int {
	return func() int {
		return 42
	}
}
