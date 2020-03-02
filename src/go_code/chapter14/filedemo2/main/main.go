package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开一个文件
	file, err := os.Open("/home/jeremyke/Desktop/b.txt")
	if err != nil {
		fmt.Println("打开失败：", err)
	}
	//关闭文件
	defer file.Close()

	//创建*Reader,是带缓冲的 默认大小4K
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err != io.EOF {                  //io.EOF文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")

}
