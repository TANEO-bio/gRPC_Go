package main

import (
	"fmt"
	"time"
)

func Wait(a int) {
	for i := 1; i <= a; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("Count: ", i)
	}
}

func Wait_channel(a int, c chan int) {
	for i := 1; i <= a; i++ {
		time.Sleep(time.Second * 1)
		c <- i
	}
	close(c)
}

func main() {
	// チャネルなし (= 返り値なし)
	start := time.Now()
	go Wait(5)
	end := time.Now()
	fmt.Println(end.Sub(start))

	// チャネルあり (= 返り値あり)
	c := make(chan int)
	start = time.Now()
	go Wait_channel(5, c)
	for d := range c {
		fmt.Println("Count: ", d)
	}
	end = time.Now()
	fmt.Println(end.Sub(start))
}
