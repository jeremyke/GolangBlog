package main

import "fmt"

func main() {
	diGui(8)
}

func diGui(n int) {
	if n > 1 {
		n--
		diGui(n)
	}
	fmt.Printf("%v", n) //1
}
