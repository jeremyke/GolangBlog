package main

import (
	"fmt"
)

//稀疏数组
func main() {
	var chessMap [11][11]int
	chessMap[1][3] = 1 //白
	chessMap[2][3] = 2 //黑

	var digitalNum int
	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			if v2 != 0 {
				digitalNum++
			}
			fmt.Printf("%d\t", v2)
		}
		fmt.Println("\t")
	}

	//转为稀疏数组
	sparseArr := make([][]int, digitalNum+1)
	sparseArrFirst := []int{len(chessMap), len(chessMap[0]), 0}
	sparseArr[0] = append(sparseArrFirst)
	n := 1
	for k1, v1 := range chessMap {
		for k2, v2 := range v1 {
			if v2 != 0 {
				sparseTmp := []int{k1 + 1, k2 + 1, v2}
				sparseArr[n] = append(sparseTmp)
				n++
			}
		}
	}

	for _, v1 := range sparseArr {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println("\t")
	}

	//转为原始数组
	row := sparseArr[0][0]
	col := sparseArr[0][1]

	var toRawArr [][]int
	for i := 1; i < row; i++ {
		toRawArr = append(toRawArr, make([]int, col))
	}

	for k3, v3 := range sparseArr {
		if k3 == 0 {
			continue
		}
		toRawArr[v3[0]-1][v3[1]-1] = v3[2]
	}

	for _, v1 := range toRawArr {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println("\t")
	}

}
