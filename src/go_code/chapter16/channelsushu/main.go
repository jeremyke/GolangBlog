package main

import (
	"fmt"
)

func readNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func getSuShu(intChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 { //不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个协程因为取不到数据就退出了")
	//这个协程取完数据就向退出管道写入数据
	exitChan <- true
}

func main() {
	//创建3个管道
	intChan := make(chan int, 8000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 4)
	//开启一个协程写入数据
	go readNum(intChan)
	//开启4个协程取出数据，并判断是否为素数
	for i := 1; i <= 4; i++ {
		go getSuShu(intChan, primeChan, exitChan)
	}
	//主线程
}
