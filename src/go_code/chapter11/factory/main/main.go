package main

import (
	"fmt"
	"go_code/chapter11/factory/model"
)

func main() {
	var stu = model.Student{
		Name:  "tom",
		Score: 79.5,
	}
	fmt.Println(stu)

	//工厂模式
	var tea = model.NewTeacher("marry", "语文")
	fmt.Println(*tea)
	fmt.Println("name=", tea.GetName(), "subject=", tea.Subject)
}
