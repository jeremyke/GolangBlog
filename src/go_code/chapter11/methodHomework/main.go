package main

import "fmt"

/**
在methodUtils结构体编写方法，从键盘接受一个整数（1-9），输出乘法表
*/

type methodUtils struct {
}

func (mu *methodUtils) getTable() {
	var num int
	fmt.Println("请输入1-9的整数")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v ", j, i, i*j)
		}
		fmt.Println()
	}
}

func (mu *methodUtils) getRes(arr [3][3]int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr[i])-i; j++ {
			tmp := arr[i][j]
			arr[i][j] = arr[j][i]
			arr[j][i] = tmp
		}
	}
	for j := 0; j < len(arr); j++ {
		for k := 0; k < len(arr[j]); k++ {
			fmt.Print(arr[j][k])
		}
		fmt.Println()
	}
}

func main() {
	var mu methodUtils
	//mu.getTable()
	var num [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num[i]); j++ {
			fmt.Print(num[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	mu.getRes(num)
}
