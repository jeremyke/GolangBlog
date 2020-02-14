package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("请输入一个正整数")
	fmt.Scanln(&n)
	res := feibolaqi(n)
	fmt.Printf("%v的菲波拉契数为%v", n, res)
}

func feibolaqi(n int) int {
	var res int
	if n == 1 || n == 2 {
		res = 1
	} else {
		res = feibolaqi(n-1) + feibolaqi(n-2)
	}
	return res
}
