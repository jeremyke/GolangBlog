package main

import (
	"fmt"
	"os"
)

func main() {
	//打开一个文件
	file, err := os.Open("/Users/jeremyke/GolandProjects/src/goStudy/src/go_code/chapter14/filedemo/test.txt")
	if err != nil {
		fmt.Println("打开失败：", err)
	}
	//输出文件
	fmt.Println(&file) //xo11121 说明：文件是一个指针
	//关闭文件
	file.Close()
}
