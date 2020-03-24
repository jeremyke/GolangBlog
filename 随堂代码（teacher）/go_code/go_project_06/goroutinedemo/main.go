package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hellow,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //s开启一个协程
	for i := 1; i <= 10; i++ {
		fmt.Println("main() hellow,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
