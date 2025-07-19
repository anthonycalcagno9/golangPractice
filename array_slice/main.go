package main

import "fmt"

func main() {
	//array literal
	a := [3]int{1, 2, 3}
	fmt.Println(a)

	//using ... will let compiler figure out how much space we need for the array
	b := [...]string{"hello", "world", "!"}
	fmt.Println(b)

	var c [2]int
	c[0] = 1
	c[1] = 2
	fmt.Println(c)

	//arrays of different lengths are different types
	fmt.Printf("type of array [3]int, a = %T\n", a)
	fmt.Printf("type of array [2]int, c = %T\n", c)

	//length function
	fmt.Println(len(a))

	//slice is a data structure with 3 elements
	//  length - len
	//  capacity - cap
	//  pointer to underlying array

	//slice literal
	xs := []string{"hello", "world"}
	fmt.Println(xs)

	//using a for range loop to iterate through the slice
	for index, value := range xs {
		fmt.Println(index)
		fmt.Println(value)
	}

	//using c style loop and accessing slice by at each element
	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i])
	}

	//append to a slice
	//append has a variadic parameter append(slice []Type, ...Type)
	xs = append(xs, "gopher", "variadic parameter")
	fmt.Printf("slice after append = %v\n", xs)

	//slicing a slice --------------------------------------------------------

	slicePractice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	//this block shows when you create a new slice based on a slice you get a copy
	//newSlice := slicePractice[0:4]
	//if you change newSlice, it will also edit that element in slicePractice as well
	fmt.Printf("Before altering slice, slicePractice = %v\n", slicePractice)
	newSlice := slicePractice[0:4]
	fmt.Printf("After slicePractice[0:4], slicePractice = %v\n", slicePractice)
	fmt.Printf("After slicePractice[0:4], newSlice = %v\n", newSlice)
	newSlice[0] = 111
	fmt.Printf("After newSlice[0] = 111, newSlice = %v\n", newSlice)
	fmt.Printf("After newSlice[0] = 111, slicePractice = %v\n", slicePractice)

	//just calling slice2[0:2] doesn't create a new slice
	//this doesn't change slice2 to become a smaller version of itself
	slice2 := []int{11, 12, 13, 14, 15, 16}
	fmt.Printf("printing slice2[0:2] = %v\n", slice2[0:2])
	fmt.Printf("printing slice2 = %v\n", slice2)

	//delete from a slice
	slice2 = append(slice2[:2], slice2[5:]...)
	fmt.Printf("slice2 after deleting from slice = %v\n", slice2)

	//using make to create a slice
	makeSlice := make([]int, 0, 10)
	fmt.Printf("makeSlice = %v\n", makeSlice)
	fmt.Println(len(makeSlice))
	fmt.Println(cap(makeSlice))

	//multidimensional slice
	anthony := []string{"anthony", "calcagno", "nachos"}
	rafa := []string{"rafa", "santos", "ice cream"}

	twoDimensionalSlice := [][]string{anthony, rafa}
	fmt.Println(twoDimensionalSlice)

	//when you want to copy a slice over to another slice you have to be careful
	//since the slice is a pointer to an underlying array, you can have multiple slices
	//pointing to the same memory address, so you will update both slices on accident
	e := []int{0, 1, 2, 3, 4, 5, 6}
	f := make([]int, 7)

	copy(f, e)

	fmt.Printf("slice f = %v\n", f)

	e[0] = 15
	fmt.Printf("slice e = %v\n", e)
	fmt.Printf("slice f = %v\n", f)

	g := []int{1, 2, 3, 4, 5, 6, 7}
	h := make([]int, 7)

	copy(h, g)

	fmt.Printf("slice g = %v\n", g)
	fmt.Printf("slice h = %v\n", h)

	g[0] = 777

	fmt.Printf("slice g = %v\n", g)
	fmt.Printf("slice h = %v\n", h)

}
