package myAccount

import (
	"fmt"
)

type FamilyAccount struct {
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	flag    bool
	detail  string
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		loop:    true,
		balance: 10000,
		flag:    false,
		detail:  "收支\t账户金额\t收支金额\t说     明\t\n",
	}
}

func (fc *FamilyAccount) ShowAccount() {
	for {
		fmt.Println("--------家庭收支记账软件-----------")
		fmt.Println("          1 收入明细")
		fmt.Println("          2 登记收入")
		fmt.Println("          3 登记支出")
		fmt.Println("          4 退出软件")
		fmt.Println("请输入（1-4）：")
		fmt.Scanln(&fc.key)
		switch fc.key {
		case "1":
			fc.details()
		case "2":
			fc.add()
		case "3":
			fc.delete()
		case "4":
			fc.exitBye()
		default:
			fmt.Println("输入有误，请重新输入")
		}
		if !fc.loop {
			fmt.Println("退出成功，BYE")
			break
		}
	}
}

//显示明细
func (fc *FamilyAccount) details() {
	fmt.Println("--------当前收支明细记录-----------")
	if fc.flag {
		fmt.Println(fc.detail)
	} else {
		fmt.Println("当前没有收支明细，来一笔吧")
	}
}

//登记收入

func (fc *FamilyAccount) add() {
	fc.flag = true
	fmt.Println("本次收入金额")
	fmt.Scanln(&fc.money)
	fc.balance += fc.money
	fmt.Println("本次收入说明")
	fmt.Scanln(&fc.note)
	fc.detail += fmt.Sprintln("收入\t", fc.balance, "\t        ", fc.money, "\t          ", fc.note, "\t\n")
}

//登记支出

func (fc *FamilyAccount) delete() {
	fmt.Println("本次支出金额")
	fmt.Scanln(&fc.money)
	if fc.money > fc.balance {
		fmt.Println("余额不足，请重试")
	} else {
		fc.flag = true
		fc.balance -= fc.money
		fmt.Println("本次支出说明")
		fmt.Scanln(&fc.note)
		fc.detail += fmt.Sprintln("支出\t", fc.balance, "\t      ", fc.money, "\t      ", fc.note, "\t\n")
	}

}

//退出

func (fc *FamilyAccount) exitBye() {
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
		fc.loop = false
	}
}
