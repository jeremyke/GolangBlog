package main

import "fmt"

func main() {
	s := "hello, world"

	fmt.Println(s[0], s[7])
	fmt.Printf("good bye" + s[:7] + "\n")

	a := "hello 世界"
	fmt.Println(len(a))

	for i, r := range "hello 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)

	}
}
