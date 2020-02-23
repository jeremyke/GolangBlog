package main

import (
	"fmt"
)

func main() {
	var arr = [5]int{23, 26, 32, 11, 4}
	maopao(&arr)
}

func maopao(arr *[5]int) {
	fmt.Println("排序前的数组为：", arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
	fmt.Println("排序后的数组为：", arr)
}
