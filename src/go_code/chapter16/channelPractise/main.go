package main

import "fmt"

//启动一个协程，将1-2000的数放入一个channel中，启动8个协程，从这个协程中取出数据，并计算1+2+...+n的值，放入宁外一个channel 最后8个协程完成后，便历结果channel
func main() {
	numChan := make(chan int, 2000)
	readEndChan := make(chan bool, 8)
	resChan := make(chan map[int]int, 2000)
	exitChan := make(chan bool, 1)

	go addNum(numChan)

	for j := 1; j <= 8; j++ {
		go readNum(numChan, readEndChan, resChan)
	}

	go readMap(resChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

func addNum(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

func readNum(numChan chan int, readEndChan chan bool, resChan chan map[int]int) {
	for {
		num, ok := <-numChan
		if !ok {
			readEndChan <- true
			if len(readEndChan) == 8 {
				close(resChan)
			}
			break
		}
		res := 0
		if num > 0 {
			for i := 1; i <= num; i++ {
				res += i
			}
		}
		aMap := make(map[int]int, 1)
		aMap[num] = res
		resChan <- aMap
	}
}

func readMap(resChan chan map[int]int, exitChan chan bool) {
	for {
		v, ok := <-resChan
		if !ok {
			break
		}
		fmt.Println(v)
	}
	close(exitChan)
}
