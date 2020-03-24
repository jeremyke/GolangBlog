package main

import (
	"fmt"
)

func TypeJudge(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("弟%v个参数是bool类型，值为 %v\n", index, v)
		case float64:
			fmt.Printf("弟%v个参数是float64类型，值为 %v\n", index, v)
		case float32:
			fmt.Printf("弟%v个参数是float32类型，值为 %v\n", index, v)
		case int:
			fmt.Printf("弟%v个参数是int类型，值为 %v\n", index, v)
		case string:
			fmt.Printf("弟%v个参数是string类型，值为 %v\n", index, v)
		default:
			fmt.Println("类型不确定或输入的类型有误\n")
			fmt.Println()
		}
	}

}

//类型断言
func main() {
	var x interface{}
	var b float64 = 1.1
	x = b //空接口可以接受任意类型
	//下面使用类型断言
	y := x.(float64)
	fmt.Printf("y的类型是 %T 值是 %v", y, y)

	//var n_1 float64 = 1.1
	//var name string = "cmh"
}
