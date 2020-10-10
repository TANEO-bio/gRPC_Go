package main

import (
	"fmt"
	"time"
)

func GetInt(a int, c chan int) {
	for i := 1; i < 5; i++ {

		c <- a + i
		fmt.Println("GetInt:", i)
	}
	close(c)
}

func main() {
	// バッファーなし
	c := make(chan int)

	go GetInt(0, c)

	for d := range c {
		time.Sleep(time.Second * 1)
		fmt.Println("Main: ", d)
	}

	// バッファーあり
	c = make(chan int, 1)

	go GetInt(0, c)

	for d := range c {
		time.Sleep(time.Second * 1)
		fmt.Println("Main: ", d)
	}
}
