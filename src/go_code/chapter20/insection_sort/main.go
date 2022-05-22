package main

import "fmt"

//倒序
func insertSort(paramArr *[]int) {
	for i := 1; i < len(*paramArr); i++ {
		for j := i; j > 0; j-- {
			if (*paramArr)[j-1] < (*paramArr)[j] {
				(*paramArr)[j-1], (*paramArr)[j] = (*paramArr)[j], (*paramArr)[j-1]
			}
		}
	}
}

func main() {
	aa := []int{7, 3, 8, 1, 9, 4, 10}
	insertSort(&aa)
	fmt.Println(aa)
}
