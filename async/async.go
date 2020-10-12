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
	// バッファーなし (= 逐次実行)
	c := make(chan int)

	go GetInt(0, c)

	for d := range c {
		time.Sleep(time.Second * 1)
		fmt.Println("Main: ", d)
	}

	// バッファーあり (= まとめて出力を得る)
	c = make(chan int, 3)

	go GetInt(0, c)

	for d := range c {
		time.Sleep(time.Second * 1)
		fmt.Println("Main: ", d)
	}
}
