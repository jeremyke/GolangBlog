package main

import "fmt"

//存放数据
func putNum(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

//取数据
func readNum(numChan chan int, resChan chan uint64, setChan chan bool) {
	for {
		num, ok := <-numChan
		if !ok {
			break
		}
		var total uint64
		for i := num; i >= 1; i-- {
			total += uint64(num)
		}
		resChan <- total
	}
	fmt.Println("有一个协程因为取不到数据就退出了")
	setChan <- true

}

func main() {
	numChan := make(chan int, 2000)
	resChan := make(chan uint64, 2000)
	setChan := make(chan bool, 8)
	go putNum(numChan)
	for i := 1; i <= 8; i++ {
		go readNum(numChan, resChan, setChan)
	}
	go func() {
		for i := 0; i < 8; i++ {
			<-setChan
		}
		close(resChan)
	}()
	for {
		num, ok := <-resChan
		if !ok {
			break
		}
		fmt.Printf("结果为%v\n", num)
	}
	fmt.Println("主线程退出")

}
