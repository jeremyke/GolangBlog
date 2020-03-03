package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行参数是：", len(os.Args))
	for k, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", k, v)
	}
}
