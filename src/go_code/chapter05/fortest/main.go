package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("请输入金字塔层数")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ { //层数
		for k := 1; k <= num-i; k++ { //空格数
			fmt.Print(" ")
		}
		for j := 1; j <= (2*i - 1); j++ { //星星数
			fmt.Print("*")
		}
		fmt.Println("\n")
	}
}
