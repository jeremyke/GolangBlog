package main

import "fmt"

func main() {
	a, b := 1, 3
	addres_a := &a
	addres_b := &b
	fmt.Printf("addres_a=%v,addres_b=%v\n", addres_a, addres_b)
	swap2(addres_a, addres_b)
	addres_a_1 := &a
	addres_b_1 := &b
	fmt.Printf("addres_a=%v,addres_b=%v\n", addres_a_1, addres_b_1)
	fmt.Printf("a=%v,b=%v\n", a, b)
}

func swap(a, b *int) {
	// a是取a的指针的值, 赋给临时变量t
	t := *a
	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t
}
func swap2(a, b *int) {
	a, b = b, a
}
