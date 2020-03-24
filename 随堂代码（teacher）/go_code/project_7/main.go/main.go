package main

import "fmt"

func main() {

	//基本数据类型在内存的布局
	var i int = 10
	//i的地址取出
	fmt.Println("i的地址 = ", &i)
	//ptr 是一个指针变量
	//ptr 的类型 *int
	//ptr 本身的值&i
	var ptr *int = &i
	fmt.Println("ptr = \n", ptr)
	fmt.Println("ptr的地址：\n", &ptr)
	fmt.Println("ptr指向的值 = ", *ptr)
}
