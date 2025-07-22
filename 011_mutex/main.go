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
	//var mu sync.Mutex // Add mutex to protect counter

	const goal = 100

	var wg sync.WaitGroup
	wg.Add(goal)

	for i := 0; i < goal; i++ {
		go func() {
			//mu.Lock() // Lock before accessing shared variable
			counter++
			//mu.Unlock() // Unlock after modifying shared variable
			//runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Num of Go routines = ", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Num of Go routines = ", runtime.NumGoroutine())
	fmt.Println("Counter = ", counter)

}
