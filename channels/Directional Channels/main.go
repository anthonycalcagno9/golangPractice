package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)

	fmt.Println("Num Go Routine Before = ", runtime.NumGoroutine())
	go foo(c)

	bar(c)
	fmt.Println("Num Go Routine End = ", runtime.NumGoroutine())

}

func foo(c chan<- int) {
	for i := 1; i < 100; i++ {
		c <- i
		fmt.Println("Num Go Routine In foo = ", runtime.NumGoroutine())

	}
	close(c)
}

func bar(c <-chan int) {
	//range keeps going as long as channel does not close
	//thats why we have to close the channel in foo()
	for v := range c {
		fmt.Println(v)
	}
}

/*
type Job struct {
	jobThatNeedsToBeDone int
}

type Worker struct {
	completedWork int
}

func main() {
	cReceiver := make(chan<- int)
	cSender := make(<-chan int)
	wg := sync.WaitGroup{}
	jobCount := 10
	workerCount := 3

	j := Job{
		jobThatNeedsToBeDone: 12,
	}

	w := Worker{
		completedWork: 0,
	}

	fmt.Printf("Type = %T\n Value = %v\n", cReceiver, cReceiver)
	fmt.Printf("Type = %T\n Value = %v\n", cSender, cSender)
	fmt.Printf("Type = %T\n Value = %v\n", wg, wg)
	fmt.Printf("Type = %T\n Value = %v\n", j, j)
	fmt.Printf("Type = %T\n Value = %v\n", w, w)

	fmt.Println("Goroutines before anything = ", runtime.NumGoroutine())
	go func(){



	}(

*/
