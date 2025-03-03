package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//create regular channel
	//CHANNELS BLOCK!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//When we write something to this channel, it will block things from moving forward until there is something ready to pull from the channel
	//Two ways to handle that, with a go routine or with a buffered channel
	c := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("go routines inside go func = ", runtime.NumGoroutine())

		c <- 42
	}()

	fmt.Println("go routines after go func = ", runtime.NumGoroutine())

	//buffered channel
	d := make(chan int, 1)

	d <- 43

	fmt.Println(<-d)

	fmt.Println(<-c)

}
