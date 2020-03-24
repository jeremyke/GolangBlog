package main

import (
	"fmt"
)

type SortInterface interface {
	sort()
}
type Sortor struct {
	name string
}

func main() {
	arry := []int{6, 1, 3, 5, 8, 4, 2, 0, 9, 7}
	learnsort := Sortor{name: "选择排序--从小到大--不稳定--n*n---"}
	learnsort.sort(arry)

	fmt.Println(learnsort.name, arry)
}

func (sorter Sortor) sort(arry []int) {
	arrylength := len(arry)
	for i := 0; i < arrylength; i++ {
		min := i
		for j := i + 1; j < arrylength; j++ {
			if arry[j] < arry[min] {
				min = j
			}
		}
		t := arry[i]
		arry[i] = arry[min]
		arry[min] = t

	}
}
