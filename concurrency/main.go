package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Waitgroup is used to handle concurrent go routines
// You create the wg, then you can use wg.Add(), wg.Done(), wg.Wait()
// You wg.add(3) to say wait for 3 go routines to end
// You use wg.Done() at the end of a goroutines code to decrement the add() count by one
// And you use wg.Wait() to show you program where to wait until wg.Done() has been called enough times
var wg sync.WaitGroup

func main() {
	fmt.Println("OS = ", runtime.GOOS)
	fmt.Println("ARCH = ", runtime.GOARCH)
	fmt.Println("Num of CPU cores = ", runtime.NumCPU())
	fmt.Println("Num of Go routines = ", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("Num of CPU cores = ", runtime.NumCPU())
	fmt.Println("Num of Go routines = ", runtime.NumGoroutine())

	wg.Wait()

}

func foo() {
	for i := 0; i < 7; i++ {
		fmt.Println("foo = ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 7; i++ {
		fmt.Println("bar = ", i)
	}
}
