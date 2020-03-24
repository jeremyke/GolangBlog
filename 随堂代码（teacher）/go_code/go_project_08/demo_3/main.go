package main

import (
	"fmt"
)

//writer Data
func writerData(intChan chan int) {
	for i := 0; i < 50; i++ {
		//放入数据
		intChan <- i
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		val, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("读到了是", val)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writerData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
