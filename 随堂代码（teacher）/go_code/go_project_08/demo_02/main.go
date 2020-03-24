package main

import (
	"fmt"
	"time"
)

//防止宕机
func sayHello() {
	for i := 0; i <= 9; i++ {
		time.Sleep(time.Second)
		fmt.Println("say hello")
	}
}
func test() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("have a mistake")
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang"
}

func main() {
	go sayHello()
	go test()

	for i := 0; i <= 9; i++ {
		time.Sleep(time.Second)
		fmt.Println("say hello main()")
	}
}
