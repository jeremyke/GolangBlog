package main

import (
	"fmt"
)

func main() {
	res := getNum(1)
	fmt.Println("和为", res)
}

func getNum(n int) int {
	var num int
	if n == 10 {
		num = 1
	} else {
		num = 2 * (getNum(n+1) + 1)
	}
	return num
}
