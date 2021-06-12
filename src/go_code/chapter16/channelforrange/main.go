package main

import (
	"fmt"
)

//func main() {
//	intChan := make(chan int,100)
//	for i :=1;i<=100;i++ {
//		intChan <- i*2
//	}
//	close(intChan)
//	for v := range intChan {
//		fmt.Printf("%v\n",v)
//	}
//}

/*
开启一个writeData协程，向管道intChan中写入50个整数，开启一个readChan协程，从管道中读取writeData写入的数据。
*/
func main() {
	intChain := make(chan int, 50)
	exitChain := make(chan bool, 1)
	go write(intChain)
	go read(intChain, exitChain)
	for {
		_, ok := <-exitChain
		if !ok {
			break
		}
	}
}

func write(c chan int) {
	for i := 1; i <= 50; i++ {
		c <- i
		fmt.Printf("写入数据为：%v\n", i)
	}
	close(c)
}

func read(c chan int, e chan bool) {
	//for v := range c {
	//	fmt.Printf("%v\n",v)
	//}
	for {
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("读到数据为%v\n", v)
	}
	e <- true
	close(e)
}
