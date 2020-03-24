package main

import (
	"fmt"
	_ "runtime"
	_ "strconv"
	"sync" //加锁
	_ "time"
)

//func main(){//计算cpu的个数和最多使用多少个
//cpuNum := runtime.NumCPU()
//fmt.Println("cpunum = ",cpuNum)

//runtime.GOMAXPROCS(cpuNum - 1)
//fmt.Println("ok!")
//}

var (
	myMap = make(map[int]int, 10)
	lock  sync.Mutex //声明全局互斥锁
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	//开启多个协程
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	lock.Lock()
	//time.Sleep(time.Second * 10)//傻方法
	for i, val := range myMap {
		fmt.Printf("map[%d] = %d", i, val)
	}
	lock.Unlock()
}
