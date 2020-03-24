package main

import (
	"fmt"
)

//基本数据类型的转换
//基本数据类型专程string类型
func main() {
	//var numb1 int = 90
	var numb2 float64 = 23.368541
	//var b bool = true
	//var str string
	//var mychar byte = 'h'
	//var str = fmt.Sprintf("%d",numb1)
	//fmt.Printf("str type %T str = ",str,str)
	var str = fmt.Sprintf("%d", numb2)
	fmt.Printf("str type %T str = ", str, str)
}
