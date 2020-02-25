package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

//speak方法绑定到Person结构体
func (p Person) speak() {
	fmt.Printf("%v是一个好人\n", p.Name)
}
func (p Person) jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println("计算结果是", res)
}
func (p Person) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println("计算结果是", res)
}
func (p Person) getTotal(n int, m int) int {
	return n + m
}

func main() {
	var xiaoming Person
	xiaoming.Name = "小明"
	xiaoming.speak()
	xiaoming.jisuan()
	xiaoming.jisuan2(89)
	total := xiaoming.getTotal(56, 32)
	fmt.Println("两数之和为", total)
}
