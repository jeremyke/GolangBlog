package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	//time.Sleep(time.Second*10)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		//time.Sleep(time.Second)
		intChan <- i
		fmt.Printf("writeData 写数据为%v\n", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		//time.Sleep(time.Second)
		fmt.Printf("readData 读到数据为%v\n", v)
	}
	//读完之后需要向exitChan写入数据true
	exitChan <- true
	close(exitChan)
}
