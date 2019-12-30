package main

import "fmt"

func coma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return coma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	news := coma("abcdefgh")
	fmt.Println(news)
}
