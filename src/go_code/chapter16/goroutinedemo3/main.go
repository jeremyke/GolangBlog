package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	totalMap = make(map[int]int, 200)
	//声明全局互质锁synchornized
	lock sync.Mutex
)

func jiecheng(n int) {
	res := 1
	for j := n; j > 0; j-- {
		res *= j
	}
	lock.Lock() //加锁
	totalMap[n] = res
	lock.Unlock() //解锁
}

func main() {
	for i := 20; i > 0; i-- {
		go jiecheng(i)
	}
	time.Sleep(time.Second * 10) //fatal error: concurrent map writes

	lock.Lock()
	for k, v := range totalMap {
		fmt.Printf("%d! = %v\n", k, v)
	}
	lock.Unlock()
}
