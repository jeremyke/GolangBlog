package main

import (
	"fmt"
)

func main() {
	str := "hello@abc"
	slice := str[6:]
	fmt.Println("slice=", slice)
	slice2 := []byte(str)
	slice2[0] = 'x'
	str2 := string(slice2)
	fmt.Println("str2", str2)
	slice3 := []rune(str)
	slice3[1] = '你'
	str3 := string(slice3)
	fmt.Println("str3", str3)

}
