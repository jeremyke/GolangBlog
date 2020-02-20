package main

import (
	"fmt"
)

func main() {
	var intArr [5]int = [...]int{1, 23, 45, 21, 46}

	aslice := intArr[1:3]
	fmt.Println(aslice)
	fmt.Println("切片的长度是：", len(aslice))
	fmt.Println("切片的容量是：", cap(aslice))
}
