//闭包
package main

import (
	"fmt"
	"strings"
)

//累加器
func AddUpper() func(int) int {
	var n int = 20
	return func(x int) int {
		n = n + x
		fmt.Println("h:", n)
		return n
	}
}

//string makeSuffix(suffix string)
//接受文件名，返回后缀
//var name string
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	g := makeSuffix(".jpg")
	fmt.Println("返回的文件名：", g("roll"))
	fmt.Println("返回的文件名：", g("foll"))
	fmt.Println("返回的文件名：", g("eoll"))
	fmt.Println("返回的文件名：", g("boll"))
	fmt.Println("返回的文件名：", g("holl"))
}
