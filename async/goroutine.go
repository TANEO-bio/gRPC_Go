package main

import (
	"fmt"
	"time"
)

func Sleep(a int) {
	println("実行中")
	time.Sleep(time.Second * 1)
}

func main() {
	// 普通の実行
	start := time.Now()
	Sleep(1)
	end := time.Now()
	fmt.Println(end.Sub(start))

	// goroutine
	start = time.Now()
	go Sleep(1)
	end = time.Now()
	fmt.Println(end.Sub(start))
}
