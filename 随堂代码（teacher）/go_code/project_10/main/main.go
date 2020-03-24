package main

import (
	"fmt"
)

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n = ", n)
}
func main() {
	test(4)
}
