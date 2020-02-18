package main

import "fmt"

//课堂练习：
func main() {
	var days int = 97
	var weeks, exdays int
	weeks = days / 7
	exdays = days % 7
	fmt.Printf("%d个星期零%d天\n", weeks, exdays)
	//2.定义一个变量保存华氏温度，华氏温度-->摄氏温度：5/9**(华氏温度-100)，请求出华氏温度对应的摄氏温度
	var hua float64 = 189.89
	var she float64 = hotconv(hua)
	fmt.Printf("华氏温度为%f转化为摄氏温度为%f", hua, she)
}

func hotconv(hua float64) float64 {
	var she float64 = (hua - 100) * 5 / 9
	return she
}
