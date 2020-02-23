package main

import (
	"fmt"
)

func main() {
	var wang = [4]string{"白眉鹰王", "金毛狮王", "青翼蝠王", "紫衫龙王"}
	var option string
	fmt.Println("请输入四大法王的名号")
	fmt.Scanln(&option)
	find(option, wang)
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

func erfen(option int, arr [5]int) {

}
