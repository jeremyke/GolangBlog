package main

import "fmt"

//倒序
func selectSort(sortArr *[]int) {
	if sortArr != nil && len(*sortArr) >= 1 {
		for i := 0; i < len(*sortArr); i++ {
			var maxNum int
			var maxNumIndex int
			for j := i; j < len(*sortArr); j++ {
				if (*sortArr)[j] >= maxNum {
					maxNum, maxNumIndex = (*sortArr)[j], j
				}
			}
			if (*sortArr)[i] < maxNum {
				(*sortArr)[i], (*sortArr)[maxNumIndex] = maxNum, (*sortArr)[i]
			}
		}
	}
}

func main() {
	aa := []int{7, 3, 8, 1, 9, 4, 10}
	selectSort(&aa)
	fmt.Println(aa)
}
