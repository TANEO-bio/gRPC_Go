package main

import (
	"fmt"
	"time"
)

func sleep(ch chan<- int, i int) {
	time.Sleep(time.Duration(i) * time.Second)
	ch <- i
}

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	defer close(c1)
	defer close(c2)

	go sleep(c1, 2)
	go sleep(c2, 3)

	for i := 0; i < 6; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		default:
			time.Sleep(1 * time.Second)
			println("default")
		}
	}
}
