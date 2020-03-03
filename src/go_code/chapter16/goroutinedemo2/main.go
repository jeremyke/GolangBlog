package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuMum := runtime.NumCPU()
	fmt.Println(cpuMum)

	//设置使用cpu个数
	runtime.GOMAXPROCS(cpuMum - 1)
	fmt.Println("ok")
}
