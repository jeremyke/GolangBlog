package main

import (
	"fmt"
)

func main() {
	var n1, n2 float64
	var operation string
	fmt.Println("请输入第一个数")
	fmt.Scanln(&n1)
	fmt.Println("请输入第二个数")
	fmt.Scanln(&n2)
	fmt.Println("请输入操作符")
	fmt.Scanln(&operation)
	var value = getValue(n1, n2, operation)
	fmt.Printf("%v %s %v的结果是%v", n1, operation, n2, value)
}

func getValue(n1 float64, n2 float64, operation string) float64 {
	var value float64
	switch operation {
	case "+":
		value = n1 + n2
	case "-":
		value = n1 - n2
	case "*":
		value = n1 * n2
	case "/":
		value = n1 / n2
	default:
		fmt.Println("输入的操作符有误")
	}
	return value
}
