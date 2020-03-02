package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "/home/jeremyke/Desktop/b.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败：", err)
	}
	defer file.Close()

	//写入缓冲流
	str := "hello klp\n"
	writer := bufio.NewWriter(file)
	for i := 5; i > 0; i-- {
		writer.WriteString(str)
	}
	//刷入文件
	writer.Flush()
}
