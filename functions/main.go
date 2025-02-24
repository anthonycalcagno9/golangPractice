package main

import "fmt"

type person struct {
	first string
}

func main() {
	foo()

	bar("anthony")

	defer otherDeferPrint()

	s := aloha("ant")
	fmt.Println(s)

	s1, n := dogYears("doge", 10)
	fmt.Println(s1, n)

	sum := sumNums(1, 2, 3, 4)
	fmt.Printf("sum = %v\n", sum)

	//defer makes it so that line runs after the function ends (naturally, due to a panic, etc)
	//defer is useful for closing down connections and cleaning up resources
	//With mutiple defers, it is LIFO. the defer that happens last, will actually be run first compared to the other defers
	defer printToExplainDefer()

	noInput := sumNums()
	fmt.Println(noInput)

	//Example showing how you can pass in a slice and unfurl it with ...
	slice1 := []int{1, 2, 3, 4, 5, 5, 6, 7}
	sum2 := sumNums(slice1...)
	fmt.Println(sum2)

	p := person{
		first: "anthony",
	}

	p.speak()

}

//func (r receiver) identifier(p parameter(s)) (return(s)) {code}

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("They call me Mr. ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}

// Variadic parameter function
// when you pass in a bunch of numbers, go treats it as []int
// variadic parameter HAS TO BE THE LAST PARAMETER in the function signature
// If you pass in a []T... as the last parameter, then that slice is used as is and no new []T needs to be created
func sumNums(nums ...int) int {
	fmt.Println(nums)
	fmt.Printf("%T\n", nums)

	n := 0
	for _, v := range nums {
		n += v
	}

	return n
}

func printToExplainDefer() {
	fmt.Println("DEFER ----------------------------------- PRINT")
}

func otherDeferPrint() {
	fmt.Println("OTHER PRINT -----------------------------------")
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}
