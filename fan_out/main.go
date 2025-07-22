package main

import (
	"fmt"
	"sync"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

}

func populate(c1 chan<- int) {
	for i := 0; i < 100; i++ {
		c1 <- i
	}
	close(c1)
}

func fanOutIn(c1 <-chan int, c2 chan<- int) {

	var wg sync.WaitGroup

	//Here were range through c1 to get each piece of work that needs to be done
	//Then we create a new go routine for each piece of work and we pass v into that func (Closure)
	//The go func will pass v, the work that needs to be done into the doWork function
	//Then once that work is done it puts that finished work back on the channel and fan's in so we get all finished work back on one channel
	//Need to wg.Add and wg.Done because we want to make sure we wait for all the go routines to finish doing their work before we close the channel
	for v := range c1 {
		wg.Add(1)
		go func(v int) {
			c2 <- doWork(v)
			wg.Done()
		}(v)
	}

	wg.Wait()
	close(c2)

}

func doWork(v int) int {
	return v * 2
}
