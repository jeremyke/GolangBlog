package main

import (
	"fmt"
	"strconv"
)

func main1() {
	var a byte = 'n'
	fmt.Printf("%d\n", a)
	fmt.Printf("%c\n", a)
	fmt.Printf("%v\n", a)
	b := fmt.Sprintf("%c\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%s", b)
}

func main2() {
	var str1 string = "hello"
	var str2 string = "123"
	int1, _ := strconv.ParseInt(str1, 10, 64)
	int2, _ := strconv.ParseInt(str2, 10, 64)
	fmt.Printf("int1's type %T int1=%v\n", int1, int1) // string hello无法转为整数，只能设置位0值
	fmt.Printf("int2's type %T int2=%v\n", int2, int2)
}
func main() {
	var n = 100
	var ptr *int = &n
	fmt.Printf("ptr's type %T ptr=%v\n", ptr, ptr)

}
