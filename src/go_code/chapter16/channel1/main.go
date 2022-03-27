package main

import (
	"fmt"
	"time"
)

func main() {
	//intChan := make(chan int,10)
	flagChan := make(chan bool, 2)
	//flagChan <- true
	//for{
	//	<- aChan
	//	//fmt.Printf("ok=%v",ok)
	//	//fmt.Printf("v=%v",v)
	//}
	//aChan <- 1
	//aChan <- 2
	//<- aChan
	//<- aChan
	//go writeChan(intChan)
	//go readChan(intChan,flagChan)
	go func() {
		time.Sleep(time.Second * 4)
		flagChan <- true
		time.Sleep(time.Second * 4)
		//close(flagChan)
	}()
	for {
		fmt.Println("OKOKOKOK")
		v, ok := <-flagChan
		fmt.Printf("v是%v，ok是%v", v, ok)
		if !ok {
			break
		}
	}

	//go func() {
	//	fmt.Println("aaaaaa")
	//	for {
	//		fmt.Println("OKOKOKOK")
	//		v,ok := <-flagChan
	//		fmt.Printf("v是%v，ok是%v",v,ok)
	//		if !ok {
	//			break
	//		}
	//	}
	//}()
	fmt.Println("999999")
	time.Sleep(time.Second * 10)
}

func writeChan(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("写入i为", i)
	}
	close(intChan)
}

func readChan(intChan chan int, flagChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("读取V是", v)
	}
	flagChan <- true
	close(flagChan)
}
