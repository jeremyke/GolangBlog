package main

import (
	"fmt"
)

func main() {
	var intArr [5]int = [...]int{100, 200, 300, 400, 500}
	slice1 := intArr[:]
	fmt.Printf("slice1的地址是：%p\n", &slice1)
	slice1 = append(slice1, 1, 2, 55, 45, 3, 5, 67, 343, 333, 67, 2)
	fmt.Printf("slice1的地址是：%p", &slice1)
}
