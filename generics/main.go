package main

import "fmt"

type typeSet interface {
	~int | ~float64
}

func main() {
	//Interesting because I thought addT(1, 1.2) would not work
	//Go compiler decided to treat 1 like a float instead of an int
	fmt.Println("adding ints using a generic = ", addT(1, 1))
	fmt.Println("adding float64s using a generic = ", addT(1.2, 1.2))
	fmt.Println("adding int + float64, shouldn't work, using a generic = ", addT(1, 1.2))

	//These functions are using a type set - An interface that has types
	fmt.Println("adding ints using a generic = ", addT2(1, 1))
	fmt.Println("adding float64s using a generic = ", addT2(1.2, 1.2))
	fmt.Println("adding int + float64, shouldn't work, using a generic = ", addT2(1, 1.2))

	//Type Alias and underlying type constraints
	type myAlias int
	var n myAlias = 42
	//We got an error here calling addT2, since we didn't have the ~ up in typeSet
	//adding that ~ means ints and also aliases with an underlying type of int
	fmt.Println("adding ints using a generic = ", addT2(n, 1))

	//This code snippet below does not work with this setup, because now the compiler
	//knows that n is a myAlias (int) so it throws and error since it's int + float64
	//fmt.Println("adding int + float64, shouldn't work, using a generic = ", addT2(n, 1.2))

	//Concrete type vs interface type
	//concrete type - means that you can instatiate an instance of that type: int, string, struct, slice
	//interface type - an interface, which is more like a contract, you never create an instance of an interface
}

// having int | float64 is called a type constraint
func addT[T int | float64](a, b T) T {
	return a + b
}

// could also have a type set, which just turns the constraint into an interface
func addT2[T typeSet](a, b T) T {
	return a + b
}
