package main

import (
	"errors"
	"fmt"
)

func test() {
	//使用defer和recover来捕获和处理异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}

func test2() {
	err := readConf("config2.ini")
	if err != nil {
		//如果读取文件错误，就输出错误，并终止程序
		panic(err)
	}
	fmt.Println("程序继续执行...")
}

func main() {
	test2()
	fmt.Println("main后面的代码...")
}
