package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("请输入一个正整数")
	fmt.Scanln(&n)
	res := getValue(n)
	fmt.Printf("%v函数值为%v", n, res)
}

func getValue(n int) int {
	var res int
	if n == 1 {
		res = 3
	} else {
		res = 2*getValue(n-1) + 1
	}
	return res
}
