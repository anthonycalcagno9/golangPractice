package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Num of CPU cores = ", runtime.NumCPU())
	fmt.Println("Num of Go routines = ", runtime.NumGoroutine())

	counter := 0

	const goal = 100

	var wg sync.WaitGroup
	wg.Add(goal)

	for i := 0; i < goal; i++ {
		go func() {
			counter++
			//runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Num of Go routines = ", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Num of Go routines = ", runtime.NumGoroutine())
	fmt.Println("Counter = ", counter)

}
