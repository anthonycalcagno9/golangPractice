package main

import "fmt"

func main() {
	x := 42

	fmt.Println(x)
	fmt.Println(&x)                // Prints the memory address of x
	fmt.Printf("%v\t%T\n", &x, &x) // Prints the type of the address, which is *int

	//defining a pointer
	y := &x                      // y is a pointer to x
	fmt.Println(y)               // Prints the memory address of x
	fmt.Printf("%v\t%T\n", y, y) // Prints the type of the pointer, which is *int

	fmt.Println(*y)                // Dereferencing y gives us the value of x, which is 42
	fmt.Printf("%v\t%T\n", *y, *y) // Prints the value of x and its type, which is int

	// Changing the value of x through the pointer y
	*y = 100
	fmt.Println(x)  // Now x is 100
	fmt.Println(*y) // Dereferencing y still gives us the updated value of x, which

	///////////////////////functions

	a := 42

	fmt.Println("a = ", a)
	deltaInt(&a)
	fmt.Println("a after pointer function = ", a)

	s := "boring old string"
	fmt.Println("s = ", s)
	deltaString(&s)
	fmt.Println("s after pointer function = ", s)

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println("xi = ", xi)
	deltaSlice(xi)
	fmt.Println("xi after pointer function = ", xi)

	m := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("m = ", m)
	deltaMap(m, "one")
	fmt.Println("m after pointer function = ", m)

}

func deltaInt(x *int) {
	*x = 777
}

func deltaString(s *string) {
	*s = "tacos"
}

// we don't need to pass a pointer here because slices are reference types
// aka the slice already is a pointer to the underlying array
func deltaSlice(s []int) {
	s[0] = 999
}

// maps are reference types too
func deltaMap(m map[string]int, s string) {
	m[s] = 1000
}
