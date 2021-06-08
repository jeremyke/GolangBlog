package main

import "fmt"

//累加器 返回的是一个函数。
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	f := AddUpper()
	fmt.Printf("f的类型是%T,f的值是%v\n", f, f)
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16
}
