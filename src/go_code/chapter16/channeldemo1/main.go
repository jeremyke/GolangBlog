package main

import (
	"fmt"
)

func main() {
	//声明
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Println(intChan) //0xc42007c080 是一个地址，引用类型
	//写入
	//往channel写入数据时不能超过容量cap！！！
	intChan <- 10
	num := 200
	intChan <- num
	//取出数据
	num2 := <-intChan
	fmt.Println(num2) //最先写入的数据被取出来了，取完了之后会报错！！！
	//d打印
	fmt.Println(len(intChan), cap(intChan))
}
