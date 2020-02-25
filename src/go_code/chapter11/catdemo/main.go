package main

import (
	"fmt"
)

//定义一个cat类，将cat类的各个字段放入cat结构体
type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 10
	cat1.Color = "白色"
	fmt.Println(cat1)
	fmt.Println(cat1.Color)
}
