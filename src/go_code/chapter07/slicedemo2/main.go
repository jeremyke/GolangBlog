package main

import (
	"fmt"
)

func main() {
	var aslice []float64 = make([]float64, 5) //对于slice必须先make再使用。var aslice []float64不行的
	fmt.Println(aslice)
	fmt.Println("切片的长度是：", len(aslice))
	fmt.Println("切片的容量是：", cap(aslice))
}
