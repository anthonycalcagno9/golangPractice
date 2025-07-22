package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	quit <- true // Signal completion
}

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("Even:", v)
		case v := <-odd:
			fmt.Println("Odd:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("Quit was false")
				return
			} else {
				fmt.Println("Quit was true:", i)
				return
			}
		}
	}
}
