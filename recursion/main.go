package main

import "fmt"

//anything you can do with recursion you can do with loops
//so use loops was the advice

func main() {
	x := factorial(5)
	fmt.Println(x)
}

func factorial(n int) int {
	fmt.Println("this is n = ", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//fac(5) -> return 5 * factorial(4)
//fac(4) -> return 4 * factorial(3)
//fac(3) -> return 3 * factorial(2)
//fac(2) -> return 2 * factorial(1)
//fac(1) -> return 1 * factorial(0)
//fac(0) -> return 1
