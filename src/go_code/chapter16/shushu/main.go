package main

import "fmt"

func main() {
	numChan := make(chan int, 2000)
	finishChan := make(chan bool, 5)
	resChan := make(chan int, 2000)

	go writeNum(numChan)
	for i := 1; i <= 4; i++ {
		go isShu(finishChan, numChan, resChan)
	}
	for {
		v, ok := <-resChan
		if !ok {
			break
		}
		fmt.Println("素数为：", v)
	}

}

func writeNum(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

func isShu(finishChan chan bool, numChan chan int, resChan chan int) {
	for {
		v, ok := <-numChan
		if !ok {
			finishChan <- true
			if len(finishChan) == 4 {
				close(resChan)
			}
			break
		}
		if v == 1 {
			continue
		} else if v == 2 {
			resChan <- v
		} else {
			flag := false
			for i := 2; i < v; i++ {
				if v%i == 0 {
					flag = true
					break
				}
			}
			if flag {
				resChan <- v
			}
		}
	}
}
