package main

import (
	"fmt"
	"go_code/chapter11/accountmodel/model"
)

func main() {
	var xiaoming model.Account
	xiaoming.SetName("小明")
	xiaoming.SetYue(1000)
	xiaoming.SetPwd([6]byte{6, 3, 5, 2, 6, 1})
	fmt.Println(xiaoming)
}
