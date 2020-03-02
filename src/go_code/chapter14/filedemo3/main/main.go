package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileName := "/home/jeremyke/Desktop/a.txt"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("读取失败：", err)
	}
	fmt.Printf("%v", string(content)) //这个方法不需要显示的打开和关闭文件
}
