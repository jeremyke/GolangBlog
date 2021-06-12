package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//建立与服务器端的链接
	conn, err := net.Dial("tcp", "192.168.31.139:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	//循环发送单行数据
	for {
		reader := bufio.NewReader(os.Stdin) //代表标准输入
		line, err1 := reader.ReadString('\n')
		if err1 != nil {
			fmt.Println("reader err=", err1)
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出来了...")
			break
		}
		_, err2 := conn.Write([]byte(line + "\n"))
		if err2 != nil {
			fmt.Println("write err=", err2)
		}
	}

}
