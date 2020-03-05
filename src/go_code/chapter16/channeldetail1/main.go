package main

func main() {
	//默认情况下，管道是双向的
	var chan1 chan int //可读可写

	var chan2 chan<- int //只写

	var chan3 <-chan int //只读
}
