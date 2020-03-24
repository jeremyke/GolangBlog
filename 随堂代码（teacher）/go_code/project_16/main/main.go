package main

import (
	"fmt"
)

func fbn(n int) []uint64 {
	fbnslice := make([]uint64, n)
	fbnslice[0] = 1
	fbnslice[1] = 1
	for i := 2; i < b; i++ {
		fbnslice[i] = fbnslice[i-1] + fbnslice[i-2]
	}
	return fbnslice
}

func main() {
	//var array [5]int = [...]int{11,22,33,44,55}
	//slice := array[1:3]//切片

	//var slices []float64 = make(float64, 4,10)
	fbnslice := fbn(10)
	fmt.Println("输出：", fbnslice)

}
