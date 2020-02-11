package main

import (
	"fmt"
)

//从控制台接收用户信息【姓名，年龄，薪水，是否通过考试】
func main() {
	var name string
	var age byte
	var sal float64
	var isPass bool
	/*fmt.Println("请输入姓名")
	fmt.Scanln(&name)//输入的值直接传给了name的地址，引用传递，会改变原来的name值
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否 通过考试")
	fmt.Scanln(&isPass)*/
	fmt.Println("请依次输入姓名，年龄，薪水，是否通过考试")
	fmt.Scanf("%v %d %f %t", &name, &age, &sal, &isPass)

}
