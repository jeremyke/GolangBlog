package main()

import(
	"fmt"
)

func main(){
	//var intArr [3]int//有默认值
	//四种数值初始化的方式
	var bumArr1[3]int = [3]int{1,2,3}
	fmt.Println("bumArr1 =",bumArr1)

	var bumArr1[3] := [3]int{1,2,3}
	fmt.Println("bumArr1 =",bumArr1)

	var bumArr2[3] = [3]int{1,2,3}
	fmt.Println("bumArr1 =",bumArr2)

	var bumArr3 = [...]int{1,2,3}
	fmt.Println("bumArr1 =",bumArr3)

	var bumArr4[3] = [...]int{0:123,1:800,2:300,3:123}//指定下表
	fmt.Println("bumArr1 =",bumArr4)


	//数组的遍历
	//for index, value := range arrat01

	//array的细节
	var array01 [3]int
	array01[0] = 1
	array01[1] = 2
	//array01[2] = 3.2//类型不匹配
}