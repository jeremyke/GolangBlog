package main

import (
	"fmt"
)

func main() {
	var aMap map[string]string
	//使用make分配内存空间
	aMap = make(map[string]string, 10) //表示可以分配10对key-value
	aMap["no1"] = "宋江"
	aMap["no2"] = "吴用"
	aMap["no1"] = "武松" //覆盖前面同样的key-value对
	fmt.Println(aMap)
}
