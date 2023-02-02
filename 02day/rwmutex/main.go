package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁
// 当一个goroutine获取读锁的时候 其他的还是能获取读锁
// 当一个goroutine 获取写锁的时候 其他的获取读锁会继续 获得写锁会等待

var (
	x    = 100
	wg   sync.WaitGroup
	lock sync.RWMutex
)

func main() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		go writex()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 100; i++ {
		go readx()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}

func readx() {
	defer wg.Done()
	lock.RLock()
	time.Sleep(time.Millisecond)
	lock.RUnlock()
}

func writex() {
	defer wg.Done()
	lock.Lock()
	x += 1
	time.Sleep(time.Millisecond * 5)
	lock.Unlock()
}

// 读写锁 RWLock()中 RLock() RUnlock() 是读锁
// lock() 与 Unlock()是写锁
