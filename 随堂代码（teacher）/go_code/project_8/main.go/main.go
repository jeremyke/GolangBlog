//键盘输入语句
//导入fmt包
//调用fmt包中的fmt.Scanln()和fmt.Scanf()
package main

import "fmt"

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入名字")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)
	fmt.Println("是否通过考试")
	fmt.Scanln(&isPass)

	fmt.Println("name is:\n age is:\n,sal is:\n,isPass is:", name, age, sal, isPass)

	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试 请用空格隔开")
	fmt.Scanf("%s\n,%d\n,%f\n,%t", &name, &age, &sal, &isPass)
}
