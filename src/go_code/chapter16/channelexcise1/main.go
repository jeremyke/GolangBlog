package main

import (
	"fmt"
)

var flag = false

//存放数据
func putNum(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

//取数据
func readNum(numChan, resChan chan int) {
	for {
		num, ok := <-numChan
		if !ok {
			break
		}
		//time.Sleep(time.Second)
		fmt.Printf("readData 读到数据为%v\n", num)
		var total int
		for i := num; i >= 1; i++ {
			total += num
		}
		resChan <- total
		if len(resChan) == 2000 {
			flag = true
			close(resChan)
		}
	}
}

func main() {
	numChan := make(chan int, 2000)
	resChan := make(chan int, 2000)
	//setChan := make(chan bool,1)
	go putNum(numChan)
	for i := 1; i <= 8; i++ {
		go readNum(numChan, resChan)
	}
	for {
		if flag {
			for v := range resChan {
				fmt.Println("v=", v)
			}
			break
		}
	}

}
