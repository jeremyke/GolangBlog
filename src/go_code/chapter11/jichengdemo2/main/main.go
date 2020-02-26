package main

import (
	"fmt"
	"go_code/chapter11/jichengdemo2/model1"
)

type Teacher struct {
	model1.People
	class int
}

func main() {
	var xiaoli Teacher
	xiaoli.Name = "小丽"
	//xiaoli.age = 12
	xiaoli.class = 11
	fmt.Println(xiaoli)
}
