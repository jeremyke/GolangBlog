package utils

import (
	"fmt"
)

type Family struct {
	quit string
	key  string
	loop bool
	//定义余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//收支的详情，当有收支时进行拼接即可
	//details string = "收支\t账户\t金额\t收支金额\t说明"
	details string
	//定义一个变量，机率是否有收支行为
	flag bool
}

func NewFamily() *Family {
	return &Family{
		key:     "",
		loop:    true,
		balance: 10000.0,
		note:    "",
		flag:    false,
		details: "收支\t账户\t金额\t收支金额\t说明",
	}
}

func (this *Family) showDetils() {
	fmt.Println("家庭收支明细记账本")
	if this.flag != true {
		fmt.Println("当前没有收入，请来存入一笔吧！嘤嘤嘤，求求你了")
	} else {
		fmt.Println(this.details)
	}
}

func (this *Family) showIncome() {
	fmt.Println("家庭收支明细支出")
	fmt.Println("本次支出金额")
	fmt.Scanln(&this.money)
	if this.balance > this.money {
		this.balance -= this.money
	} else {
		fmt.Println("你没这么多钱")
	}
	fmt.Println("说明")
	fmt.Scanln(&this.note)
	//将收入情况拼接到details
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
}

func (this *Family) showPay() {
	fmt.Println("家庭收支明细收入")
	fmt.Println("本次收入金额")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("说明")
	fmt.Scanln(&this.note)
	//将收入情况拼接到details
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *Family) showBreak() {
	fmt.Println("确定要退出吗？")
	for {
		fmt.Scanln(&this.quit)
		if this.quit == "是" {
			break
		}
		fmt.Println("请重新输入")
	}
	if this.quit == "是" {
		this.loop = false
	}
}

func (this *Family) MainMenu() {
	for {
		fmt.Println("家庭收支明细记账本")
		fmt.Println("家庭收支明细支出")
		fmt.Println("家庭收支明细收入")
		fmt.Println("家庭收支明细退出")
		fmt.Println("家庭收支明细请选择（1-4）：")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showBreak()
		case "2":
			this.showDetils()
		case "3":
			this.showIncome()
		case "4":
			this.showPay()
		default:
			fmt.Println("请输入正确选项")

			if this.loop == false {
				break
			}
		}
		fmt.Println("感谢使用")
	}
}
