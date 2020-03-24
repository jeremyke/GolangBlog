package main

import (
	"fmt"
)

func main() {
	//创建管道
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("管道有点啥：%v", intChan)

	//写入数据
	intChan <- 10
	num := 56
	intChan <- num
	fmt.Printf("管道现在有点啥：%v", intChan)
	//从管道中读取数据
	//n_1 := <- intChan
}
