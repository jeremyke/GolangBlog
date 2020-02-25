package main

import "fmt"

type A struct {
	Num int
}
type B struct {
	Num int
}

type Student struct {
	score float64
}

type Stu Student

func main() {
	var a A
	var b B
	a = A(b) //可以转换，需保证结构体字段完全一样
	fmt.Println(a, b)
	var stu1 Student
	var stu2 Stu
	stu2 = stu1 //不行，二者类型不一样，可以写stu2 = Stu(stu1)

}
