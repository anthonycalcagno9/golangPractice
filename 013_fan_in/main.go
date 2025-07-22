package main

import (
	"fmt"
	"sync"
)

func main() {

	//send even and odd numbers to channels

	//received even and odd numbers to channels and then pull both of those into a fanin channel

	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Printf("fanin values = %v", v)
	}

}

func send(even, odd chan<- int) {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(even)
	close(odd)
}

func receive(even, odd <-chan int, fanin chan<- int) {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)

}
