package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan <- i * 2
	}
	//没有关闭管道会报错：fatal error: all goroutines are asleep - deadlock!
	close(intChan)
	for v := range intChan {
		fmt.Println("v=", v)
	}
}
