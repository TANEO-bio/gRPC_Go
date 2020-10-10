package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 排他制御なし
	c := 0
	for i := 0; i < 1000; i++ {
		go func() {
			c++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(c)

	// 排他制御あり
	var mu sync.Mutex
	c = 0
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()         // 排他ロック取得
			defer mu.Unlock() // 関数終了時に排他ロック解除
			c++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(c)
}
