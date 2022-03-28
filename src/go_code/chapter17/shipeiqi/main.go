package main

import (
	"fmt"
	"reflect"
)

func main() {
	tes1 := func(v1, v2 int) {
		fmt.Println("两数之和为：", v1+v2)
	}
	tes2 := func(v1 int, v2 int, s string) {
		fmt.Printf("%s的身高为%v，年龄为%v", s, v1, v2)
	}

	bridge := func(call interface{}, param ...interface{}) {
		//参数
		n := len(param)
		var inValue []reflect.Value
		for i := 0; i < n; i++ {
			inValue = append(inValue, reflect.ValueOf(param[i]))
		}
		//方法
		inFun := reflect.ValueOf(call)
		//调用
		inFun.Call(inValue)
	}
	bridge(tes1, 12, 23)
	bridge(tes2, 12, 23, "keke")
}
