package main

import (
	"fmt"
	"strconv"
	"sync"
)

// go语言内置的map不是并发安全的
var m = make(map[string]int)

// sync.Map不需要使用make 初始化了
// sync.Map 使用一些特点的方法取值 设置值
// store() 设置值 load()加载值
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v  v=:%v\n", key, get(key))
			wg.Done()
		}(i)
	}

	wg.Wait()
	UseSM()
}

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

// 这里要注意的是 sync.Map 有自己的增删改查的方法
// store() 设置值
// load() 取值
// loadOrStore() 取值或者设置值
// delete() 删除值
var m1 = sync.Map{}

func UseSM() {
	wg := sync.WaitGroup{}
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m1.Store(key, n)
			value, _ := m1.Load(key)
			fmt.Println(value)
		}(i)
	}
}
