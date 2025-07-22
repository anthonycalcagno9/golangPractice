package main

import "fmt"

// Value semantics - we pass a copy of the value into the function
// So a new copy of x is being incremented by one, but the original x will remain unchanged
func addOne(x int) {
	x++
	fmt.Println("Inside addOne:", x) // This will print the incremented value
}

// Pointer semantics - we pass a pointer to the value into the function
// So the original x is being incremented by one, and the change will be reflected outside
func addOnePointer(x *int) {
	*x++
	fmt.Println("Inside addOnePointer:", *x) // This will print the incremented value
}

func main() {
	a := 5
	fmt.Println("Original a:", a)   // a = 5
	addOne(a)                       // Pass by value
	fmt.Println("After addOne:", a) // a remains 5

	addOnePointer(&a)                      // Pass by reference
	fmt.Println("After addOnePointer:", a) // a is now 6
}

// Pointer value semantics heuristics: "Rules of thumb"
// 1. Use Value semantics when possible for simplicity and safety.
// They do not have a shared state
// If you don't need to modify input data or you're working with small data (built in types or small structs), use value semantics.

// 2. Use Pointer semantics for large data
// Copying large structs or arrays can be expensive in terms of performance.
// If 64 bytes or larger, use pointer semantics

// 3. Use Pointer semantics for mutability
// If a function needs to modify it's RECEIVER or its input data, use pointer semantics.
// Common use case: methods that need to update the state of the struct

// 4. Consistency
// Typically, once a type has methods that use pointer semantics, all methods should use pointer semantics.
