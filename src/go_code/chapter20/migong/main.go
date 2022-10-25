package main

import "fmt"

//定义一个二维数组模拟迷宫
var (
	sheet = [][]bool{
		[]bool{true, true, true, true, true},
		[]bool{true, true, true, true, true},
		[]bool{false, false, true, true, true},
		[]bool{true, true, true, true, true},
		[]bool{true, true, true, true, true},
		[]bool{true, true, true, true, true},
	}
)

func main() {
	process(sheet, []int{0, 0}, []int{5, 4}, [][]int{[]int{0, 0}})
}

//走迷宫（sheet 迷宫，start 开始坐标）
func process(sheet [][]bool, start, end []int, trace [][]int) {
	if end[0] == start[0] && end[1] == start[1] {
		fmt.Println(trace)
		return
	}
	//右走
	if start[1]+1 < len(sheet[start[0]]) && sheet[start[0]][start[1]+1] && isValid([]int{start[0], start[1] + 1}, trace) {
		//traceTmp := trace
		trace = append(trace, []int{start[0], start[1] + 1})
		process(sheet, []int{start[0], start[1] + 1}, end, trace)
		trace = trace[0 : len(trace)-1]
	}
	//下走
	if start[0]+1 < len(sheet) && sheet[start[0]+1][start[1]] && isValid([]int{start[0] + 1, start[1]}, trace) {
		trace = append(trace, []int{start[0] + 1, start[1]})
		process(sheet, []int{start[0] + 1, start[1]}, end, trace)
		trace = trace[0 : len(trace)-1]
	}
	//上走
	if start[0]-1 >= 0 && sheet[start[0]-1][start[1]] && isValid([]int{start[0] - 1, start[1]}, trace) {
		trace = append(trace, []int{start[0] - 1, start[1]})
		process(sheet, []int{start[0] - 1, start[1]}, end, trace)
		trace = trace[0 : len(trace)-1]
	}

	//左走
	if start[1]-1 >= 0 && sheet[start[0]][start[1]-1] && isValid([]int{start[0], start[1] - 1}, trace) {
		trace = append(trace, []int{start[0], start[1] - 1})
		process(sheet, []int{start[0], start[1] - 1}, end, trace)
		trace = trace[0 : len(trace)-1]
	}
}

//判断是否走过的格子
func isValid(point []int, trace [][]int) bool {
	flag := true
	for _, arrVal := range trace {
		if arrVal[0] == point[0] && arrVal[1] == point[1] {
			flag = false
			break
		}
	}
	return flag
}
