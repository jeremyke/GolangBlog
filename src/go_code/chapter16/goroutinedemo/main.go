package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Printf("hello world " + strconv.Itoa(i) + "\n")
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //开启一个协程执行test函数
	for i := 0; i < 10; i++ {
		fmt.Printf("hello golang " + strconv.Itoa(i) + "\n")
		time.Sleep(time.Second)
	}
}
