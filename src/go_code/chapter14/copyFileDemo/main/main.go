package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := CopyFile("/home/jeremyke/Desktop/bb.png", "/home/jeremyke/Desktop/aa.png")
	if err != nil {
		fmt.Println("拷贝失败：", err)
	} else {
		fmt.Println("拷贝成功")
	}
}

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	//读取文件
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("文件读取失败:", err)
	}
	defer srcFile.Close()
	read := bufio.NewReader(srcFile)
	//写入文件
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件写入失败：", err)
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	//调用copy
	return io.Copy(writer, read)
}
