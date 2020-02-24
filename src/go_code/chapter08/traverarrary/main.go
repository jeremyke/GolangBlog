package main

import (
	"fmt"
)

func main() {
	var arr = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	//for循环遍历
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}
	//for-range遍历
	for _, value := range arr {
		for _, value1 := range value {
			fmt.Print(value1)
		}
		fmt.Println()
	}
}
