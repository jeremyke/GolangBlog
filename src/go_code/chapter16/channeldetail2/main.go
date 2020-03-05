package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Printf("hello%v", i)
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出错了:", err)
		}
	}()
	var a map[int]string
	a[0] = "sddddd"
}

func main() {
	go sayHello()
	go test()
	time.Sleep(time.Second)
}
