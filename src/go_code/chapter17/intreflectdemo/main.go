package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {
	//获取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType) //int
	//获取reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue=%v 类型为：%T\n", rValue, rValue) //100 reflect.Value
	//将reflect.Value-->interface{}
	iv := rValue.Interface()
	//interface{}-->原来的类型
	num2 := iv.(int)
	fmt.Println("num2=", num2)
}

func main() {
	num := 100
	reflectTest(num)
}
