package main

import "fmt"

func main() {
	x := 1
	p := &x
	c := *p
	fmt.Println(x)  //1
	fmt.Println(&x) //变量x的地址 0xc042050080
	fmt.Println(*p) //x的地址对应的值 1
	fmt.Println(*(&x))
	fmt.Println(c)
}
