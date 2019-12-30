package main

import (
	"fmt"
	"strings"
)

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i-1]
			break
		}
	}
	return s
}
func basename2(s string) string {
	slast := strings.LastIndex(s, "/")
	s = s[slast+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	news := basename2("aa/dd/name.go")

	fmt.Println(news)
}
