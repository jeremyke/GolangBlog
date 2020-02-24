package main

import (
	"fmt"
)

func main() {
	/*var wang = [4]string{"白眉鹰王", "金毛狮王", "青翼蝠王", "紫衫龙王"}
	var option string
	fmt.Println("请输入四大法王的名号")
	fmt.Scanln(&option)
	find(option, wang)*/

	arr := [5]int{1, 6, 23, 46, 88}
	erfen(&arr, 6, 0, 4)
}

func find(option string, wang [4]string) {
	for i := 0; i < len(wang); i++ {
		if wang[i] == option {
			fmt.Printf("查找到了该名号，其index为%v\n", i)
			break
		} else if i == (len(wang) - 1) {
			fmt.Printf("没有找到%v", option)
		}
	}
}

//二分查找
func erfen(arr *[5]int, option int, leftIndex int, rightIndex int) {
	if leftIndex > rightIndex {
		fmt.Println("不存在找不到")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > option {
		erfen(arr, option, leftIndex, middle-1)
	} else if (*arr)[middle] < option {
		erfen(arr, option, middle+1, rightIndex)
	} else {
		fmt.Println("找到了，下标为：", middle)
	}
}
