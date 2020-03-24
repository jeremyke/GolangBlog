package main

import (
	"fmt"
)

type Calcuator struct {
	Num_1 float64
	Num_2 float64
}

func (calcuator *Calcuator) getSum() float64 {
	return calcuator.Num_1 + calcuator.Num_2
}

type Person struct {
	Name string
}

type MethodUtils struct {
}

func (mu MethodUtils) print_1() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) print_2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Println("*")
		}
		fmt.Println()
	}
}

type Circle struct {
	radius float64
}

type integer int //定义自己的方法

func (i integer) print() {
	fmt.Println("123", i)
}

func (i *integer) charge() {
	*i = *i + 1
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

//给Person类型绑定一种方法
func (p Person) test() {
	fmt.Println("test()", p.Name)
}

func (p Person) speak() {
	fmt.Println(p.Name, "is a good man")
}

func (p Person) jisuna() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(p.Name, "the anower is", sum)
}

func (calcuator *Calcuator) getRes(operator byte) float64 { //计算器
	res := 0.0
	switch operator {
	case '+':
		res = calcuator.Num_1 + calcuator.Num_2
	case '-':
		res = calcuator.Num_1 - calcuator.Num_2
	case '*':
		res = calcuator.Num_1 * calcuator.Num_2
	case '/':
		res = calcuator.Num_1 / calcuator.Num_2
	default:
		fmt.Println("输入有误!")
	}
	return res
}

func (p Person) jisuan_1(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println("the anower is~~~~~", sum)
}

func main() {
	var p Person
	p.test() //调用方法
	p.speak()
	p.jisuna()
	var a int
	fmt.Println("请输入所要累加的值")
	fmt.Scanf("%d", &a)
	p.jisuan_1(a)

	var c Circle
	c.radius = 4.0
	res := c.area()
	fmt.Println("面积是：", res)

	var i_1 integer
	i_1 = 10
	i_1.print()

	var mu MethodUtils
	mu.print_1()
	var mu_1 MethodUtils
	mu_1.print_2(10, 20)
}
