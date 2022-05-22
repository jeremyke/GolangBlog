package main

import "fmt"

func bubble(sortArr *[]int) {
	if sortArr != nil && len(*sortArr) >= 1 {
		for k := len(*sortArr) - 1; k > 0; k-- {
			for i := k; i > 0; i-- {
				if (*sortArr)[i] < (*sortArr)[i-1] {
					(*sortArr)[i], (*sortArr)[i-1] = (*sortArr)[i-1], (*sortArr)[i]
				}
			}
		}
	}
}

func main() {
	aa := []int{7, 3, 8, 1, 9, 4, 10}
	bubble(&aa)
	fmt.Println(aa)
}
