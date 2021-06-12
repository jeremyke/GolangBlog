package main

import (
	"fmt"
	"time"
)

func readNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
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
	intChan := make(chan int, 10000)
	primeChan := make(chan int, 10000)
	exitChan := make(chan bool, 8)
	//	start := time.Now().Unix()
	//开启一个协程写入数据
	go readNum(intChan)
	//开启4个协程取出数据，并判断是否为素数
	for i := 1; i <= 8; i++ {
		go getSuShu(intChan, primeChan, exitChan)
	}
	//主线程在退出管道里面取出4个true，说明素数管道处理完毕
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Printf("使用协程耗费时间：%v秒\n", end-start)
		close(primeChan)
	}()
	//遍历素数管道
	for {
		num, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数为%v\n", num)
	}
	fmt.Println("主线程退出")

}
