package main

import (
	"fmt"
)

func main() {
	var key string
	loop := true
	balance := 10000.0
	var money float64
	var note string
	flag := false
	detail := "收支\t账户金额\t收支金额\t说     明\t\n"
	for {
		fmt.Println("--------家庭收支记账软件-----------")
		fmt.Println("          1 收入明细")
		fmt.Println("          2 登记收入")
		fmt.Println("          3 登记支出")
		fmt.Println("          4 退出软件")
		fmt.Println("请输入（1-4）：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("--------当前收支明细记录-----------")
			if flag {
				fmt.Println(detail)
			} else {
				fmt.Println("当前没有收支明细，来一笔吧")
			}
		case "2":
			flag = true
			fmt.Println("本次收入金额")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明")
			fmt.Scanln(&note)
			detail += fmt.Sprintln("收入\t", balance, "\t        ", money, "\t          ", note, "\t\n")
		case "3":
			fmt.Println("本次支出金额")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足，请重试")
				break
			}
			flag = true
			balance -= money
			fmt.Println("本次支出说明")
			fmt.Scanln(&note)
			detail += fmt.Sprintln("支出\t", balance, "\t      ", money, "\t      ", note, "\t\n")
		case "4":
			fmt.Println("确定要退出吗？y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("输入有误，请重新输入")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("输入有误，请重新输入")
		}
		if !loop {
			fmt.Println("退出成功，BYE")
			break
		}
	}
}
