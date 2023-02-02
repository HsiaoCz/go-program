package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作
// 这里不加锁 会使得每次得到的结果都不同
// 加锁可以避免这种情况
// 原子操作 使对数据的修改是原子的
// atomic func
// load readfunc 返回对应的值，第一个参数是指针
// store 写入操作，第一个参数是指针，第二个参数是要写入的值
// add 修改操作，第一个要修改的数据的指针，第二个修改的值
// swap 交换操作
// compareAndSwap 比较并交换值
var x int64
var wg sync.WaitGroup

//var lock sync.Mutex

func add() {
	//lock.Lock()
	//x++
	//lock.Unlock(
	//使用原子操作传递指针
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
