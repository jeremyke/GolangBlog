package main

import (
	"flag"
	"fmt"
)

func main() {
	var user, pwd, host string
	var port int
	//这种方法参数顺序可以随意
	flag.StringVar(&user, "u", "", "用户名")
	flag.StringVar(&pwd, "pwd", "", "密码")
	flag.StringVar(&host, "h", "localhost", "主机名")
	flag.IntVar(&port, "port", 3306, "端口号")

	flag.Parse() //转换，必须调用该方法。

	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}
