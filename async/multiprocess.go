package main

import (
	"fmt"
	"sync"
	"time"
)

func sleep(sec time.Duration) {
	time.Sleep(time.Second * sec)
}

func wgSleep(wg *sync.WaitGroup, sec time.Duration) {
	time.Sleep(time.Second * sec)
	wg.Done()
}

func main() {
  // sync.WaitGroupで並列処理のタスクを終了待ち
	var wg sync.WaitGroup

  // 単一スレッド
	start := time.Now()
	sleep(1)
	sleep(2)
	end := time.Now()
	fmt.Println(end.Sub(start)) // 3s

  // 並列処理
	start = time.Now()
	wg.Add(2)
	go wgSleep(&wg, 1)
	go wgSleep(&wg, 2)
	wg.Wait()
	end = time.Now()
	fmt.Println(end.Sub(start)) // 2s
}
