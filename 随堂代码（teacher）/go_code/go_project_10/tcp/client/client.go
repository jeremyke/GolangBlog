package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tap", "127.0.0.1 : 8888")
	if err != nil {
		fmt.Println("错误 :", err)
		return
	}
	//fmt.Println("成功 :", conn)
	//准备完成的功能
	//发送单行数据，之后退出
	reader := bufio.NewReader(os.Stdin) //os.Sdtin表示标准输入，一般是键盘【终端】
	//从终端读取一行用户输入，并准备发给服务器
	line, err := reader.ReadString('\n') //每一个空格接受
	if err != nil {
		fmt.Println("错误 :", err)
	}
	line = strings.Trim(line, " \r\n")
	if line == "exit" {
		fmt.Println("客户端退出..\n")
		return
	}
	//将line发给服务器
	_, err = conn.Write([]byte(line + "\n"))
	if err != nil {
		fmt.Println("conn.Write err =", err)
	}
	//fmt.Printf("客户端发送了 %d 字节的数据，并退出", n)
}
