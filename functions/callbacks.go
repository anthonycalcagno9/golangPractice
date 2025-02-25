package main

import "fmt"

// Callback just means passing in a function as an argument to another function
func callbacksPractice() {
	fmt.Println("hello world")

	fmt.Println(doMath(7, 9, add))
	fmt.Println(doMath(7, 9, subtract))

	//Closure practice
	f := incrementor()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 4

	g := incrementor()
	fmt.Println(g()) // 1
	fmt.Println(g()) // 2
	fmt.Println(g()) // 3
	fmt.Println(g()) // 4

}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

// Func to show closure
// the inner function is inclosed in the outer function
// so the x that is declared in the outer function is kept in the scope of the inner function
// so when we call it a couple of times in the code above, x will continue to increment by 1
// also notice how if we call incrementor twice, assigning it to f and g, those x variables
// are seperate, and will start back from 0
func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
