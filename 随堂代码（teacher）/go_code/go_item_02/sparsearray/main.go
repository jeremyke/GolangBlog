package main

import (
	"fmt"
)

type Node struct {
	row int
	col int
	val int
}

func main() {
	//创建一个数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	for i, v_1 := range chessMap {
		for j, v_2 := range v_1 {
			fmt.Printf("%d\t", v_2)
		}
		fmt.Println()
	}
	//转成稀疏数组
	var sparseArr []ValNode
	//标准的一个稀疏数组应该还有一个表示记录元素的二维数组的规模（行和列，默认值）
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	for i, v_1 := range chessMap {
		for j, v_2 := range v_1 {
			if v_2 != 0 {
				//创建一个val值节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v_2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}

	}
	//输出
	for i, valnode := range sparseArr {
		fmt.Printf("%d :  %d :  %d :  %d :  \n", i, valNode.row, valNode.col, valNode.val)
	}

	//将这个稀疏数组存盘

	//如何恢复原始数组

	//打开存盘文件

	//先创建一个原始数组
	var chessMap_2 [11][11]int
	for i, valNode := range sparseArr {
		if i != 0 { //跳过第一行数据
			chessMap_2[valNode.row][valNode.col] = valNode.val

		}
	}
	//遍历稀疏数组
	for _, valNode := range sparseArr { //这里也可以是文件路径{
		chessMap_2[valNode.row][valNode.col] = valNode.val
	}

}
