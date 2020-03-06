package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {
	rValue := reflect.ValueOf(b)
	//rValue.SetInt(20)//报错,因为此时传过来的是地址，所以rValue是个指针
	rValue.Elem().SetInt(20)
	fmt.Printf("rValue=%v", rValue.Elem().Int())

}
func main() {
	var num int = 100
	reflectTest(&num)
}
