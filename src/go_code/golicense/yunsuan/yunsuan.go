package main

import "fmt"

var x int8 = -32

func main() {
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", x>>2)
	fmt.Println(x >> 2)

	y := 1e100
	i := int(y)
	fmt.Println(y)
	fmt.Println(i)
}
