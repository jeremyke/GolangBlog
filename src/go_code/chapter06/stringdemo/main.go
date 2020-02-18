package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello我"
	fmt.Println("str len=", len(str))

	str1 := "go golang"
	fmt.Printf("index=%v\n", strings.LastIndex(str1, "go"))

	str2 := " sdfds sdfd dsfdg  "
	fmt.Printf("处理后字符串为%q\n", strings.TrimSpace(str2))

	str3 := "! he！llo! "
	fmt.Printf("处理后字符串为%q", strings.Trim(str3, "! "))
}
