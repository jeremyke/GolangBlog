package main

import (
	"fmt"
)

type A struct {
	Num int
}

func (a A) test() {
	a.Num = 11
	fmt.Println(a.Num)
}
func main() {
	var t A
	t.test()
}
