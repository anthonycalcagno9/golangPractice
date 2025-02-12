package main

import (
	"fmt"

	"math/rand/v2"
)

// Can use this type of variable declaring inside or outside of a function
var x int = 7

func main() {
	//will print variable value with %v and variable type with %T
	fmt.Printf("Printing variable value = %v and type = %T\n", x, x)

	//short hand to declare a variable, can only be used inside of a function
	y := 10

	//example if, else if, else statement
	if y < 9 {
		fmt.Printf("y is less than 9\n")
	} else if y == 9 {
		fmt.Printf("y is equal to 9\n")
	} else {
		fmt.Printf("y is greater than 9\n")
	}

	//Showing use of && the AND operator
	if y == 10 && x == 7 {
		fmt.Println("And operator is being used.")
	}

	//Showing use of || the OR operator
	if y == 10 || x == 13 {
		fmt.Println("Or operator is being used")
	}

	//Statement; statement idiom
	//Useful because it limits variable z's scope to these 5 lines
	if z := rand.Int(); z >= x {
		fmt.Printf("z is %v and that is greater than or equal to x which is %v\n", z, x)
	} else {
		fmt.Printf("condition isn't true\n")
	}

	myMap := map[string]int{
		"tacos":    5,
		"burritos": 2,
		"salsa":    7,
	}

	//Comma ok idiom
	//if value in myMap exists, tacos will be set correctly and ok will be true
	//if it doesn't exist in myMap, tacos will be set to zero value and ok will be false
	if tacos, ok := myMap["tacos"]; ok {
		fmt.Printf("was true, tacos = %v\n", tacos)
	} else {
		fmt.Printf("was false, tacos = %v\n", tacos)
	}

	//check for presence in map using blank identifier _
	_, present := myMap["burritos"]
	fmt.Printf("does value exist in the map? present = %v", present)

	//switch, case, default
	switch x {
	case 40:
		fmt.Printf("x is 40")
	case 41:
		fmt.Printf("x is 41")
	default:
		fmt.Println("switch case default print out")
	}

	//c style for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("for loop counter = %v\n", i)
	}

	j := 5
	//for that behaves like a while loop
	for j < 10 {
		fmt.Printf("while loop print j = %v\n", j)
		j++
	}

	//Can use break and continue in for loops

	//for range loop, use with slices or maps
	sliceExample := []int{6, 7, 8, 9, 10}
	for index, value := range sliceExample {
		fmt.Printf("printing for range loop index = %v and value = %v\n", index, value)
	}

	mapExample := map[string]int{
		"cat":  3,
		"dog":  5,
		"bird": 7,
	}
	for key, value := range mapExample {
		fmt.Printf("printing map for range key = %v and value = %v\n", key, value)
	}

	//using modulus to find remainder
	z := 83 % 40
	fmt.Printf("remainder of 83 / 40 should be 3. z = %v\n", z)

}
