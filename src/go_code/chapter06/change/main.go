package main

import (
	"fmt"
)

func change(a, b *int) {
	tmp := *b
	*b = *a
	*a = tmp
}

func main() {
	a := 10
	b := 20
	change(&a, &b)
	fmt.Printf("a=%d b=%d", a, b)
}
