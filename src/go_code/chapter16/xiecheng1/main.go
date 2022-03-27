package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	res  = make(map[int]int64, 20)
	lock sync.Mutex
)

//计算1-200的阶乘并且把各个数的阶乘放入map中，然后打印出来
func main() {
	for n := 1; n <= 20; n++ {
		lock.Lock()
		go getJc(n)
		lock.Unlock()
	}
	time.Sleep(time.Second * 5)
	fmt.Println(res)
}

func getJc(num int) {
	var re int64
	re = 1
	for i := int64(num); i > 0; i-- {
		re *= i
	}
	res[num] = re
}
