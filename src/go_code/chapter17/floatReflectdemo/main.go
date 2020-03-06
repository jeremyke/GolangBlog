package main

import (
	"fmt"
	"reflect"
)

func test(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType类型是", rType)
	rValue := reflect.ValueOf(b)
	fmt.Println("rValue类型是", rValue)
	aType := rValue.Type()
	fmt.Println("aType类型是", aType)
	aKind := rValue.Kind()
	fmt.Println("aKind类型是", aKind)
	aaa := rValue.Interface()
	bb := aaa.(float64)
	fmt.Println("bb=", bb)
}

func main() {
	var v float64 = 1.2
	test(v)
}
