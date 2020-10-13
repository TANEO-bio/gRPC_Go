package main

import (
	"fmt"
)

func countTen(counter chan<- int, reset <-chan bool) {
	defer close(counter)
	for i := 0; i < 10; i++ {
		select {
		case s := <-reset:
			if s {
				i = 0
				counter <- i
			} else {
				return
			}
		default:
			counter <- i

		}
	}
}

func main() {
	var (
		counter = make(chan int)
		reset   = make(chan bool, 1)
		flag    = false
	)
	defer close(reset)
	go countTen(counter, reset)

	for {
		val, ok := <-counter
		if ok == false {
			break
		}

		if flag == false && val == 3 {
			reset <- true
			flag = true
		}

		if flag == true && val == 8 {
			reset <- false
		}

		fmt.Println(val)

	}
}
