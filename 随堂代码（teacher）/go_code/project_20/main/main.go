package main

import (
	"fmt"
)

type Atm struct {
	card_1 map[string]string //map中包括卡号，金额
	card_4 int               //存款
	card_2 int               //取款
	card_3 int               //余额
	card_5 int               //转账
}

func main() {
	var arrary [100]int
	var num int
	fmt.Println("请您选择要进行的操作(1:申请新卡  2:存款   3:取款  4:查余额  5:转账 6:修改密码 0:结束循环)")
	fmt.Scanln(&num)
	if num == 1 {
		for i := 0; i < 100; i++ {
			if arrary[i] == 0 {
				//arrary[i] == int i
				fmt.Println("恭喜您成功申请新卡：", arrary[i])
			}
		}
	}
	if num == 2 {
		var cardNumb int
		var moneyNumb int
		fmt.Println("请您输入卡号：")
		fmt.Scanf("%d", &cardNumb)
		fmt.Println("请您输入索要存入的钱数：")
		fmt.Scanf("%d", &moneyNumb)
		BinaryFind(&arrary, 0, (len(arrary) - 1), cardNumb, moneyNumb)
		//caRd.card_1 = make(map[string]string)
	}
	if num == 3 {
		var cardNumb_1 int
		var moneyNumb_1 int
		fmt.Println("请您输入卡号：")
		fmt.Scanf("%d", &cardNumb_1)
		fmt.Println("请您输入所要取出的钱数：")
		fmt.Scanf("%d", &moneyNumb_1)
		BinaryFind_1(&arrary, 0, (len(arrary) - 1), cardNumb_1, moneyNumb_1)
	}
	if num == 4 {
		var cardNumb_2 int
		fmt.Println("请您输入卡号：")
		fmt.Scanf("%d", &cardNumb_2)
		BinaryFind_2(&arrary, 0, (len(arrary) - 1), cardNumb_2)

	}
	if num == 5 {
		var cardNumber_1 int
		var cardNumber_2 int
		var moneyNumber_1 int
		fmt.Println("请您输入卡号：")
		fmt.Scanf("%d", &cardNumber_1)
		fmt.Println("请您输入所要转钱的卡号：")
		fmt.Scanf("%d", &cardNumber_2)
		fmt.Println("请您输入所要转钱的金额：")
		fmt.Scanf("%d", &moneyNumber_1)
		BinaryFind_3(&arrary, 0, (len(arrary) - 1), cardNumber_2, moneyNumber_1) //本人思考要是本人账户不存在就没有转账这一可能性了，故将转账功能只查询被转账用户，减少查询时间
	}
	if num == 6 {
		fmt.Println("感谢您使用本程序")
		return
	}
}

func BinaryFind(arrary *[100]int, leftIndex int, rightIndex int, findVal int, money int) { //存款函数
	middle := (leftIndex + rightIndex) / 2

	if leftIndex > rightIndex {
		fmt.Printf("没有找到")
		return
	}

	if (*arrary)[middle] > findVal {
		BinaryFind(arrary, leftIndex, middle-1, findVal, money)
	} else {
		fmt.Println("找到了是：", arrary[middle])
		var caRd Atm
		caRd.card_4 = arrary[middle] + money
		fmt.Println("现在的余额为：", caRd.card_4)
	}
}

func BinaryFind_1(arrary *[100]int, leftIndex int, rightIndex int, findVal int, money_1 int) { //提现函数
	middle := (leftIndex + rightIndex) / 2

	if leftIndex > rightIndex {
		fmt.Printf("没有找到")
		return
	}

	if (*arrary)[middle] > findVal {
		BinaryFind(arrary, leftIndex, middle-1, findVal, money_1)
	} else {
		fmt.Println("找到了是：", arrary[middle])
		var caRd Atm
		caRd.card_4 = arrary[middle] - money_1
		fmt.Println("现在的余额为：", caRd.card_4)
	}
}

func BinaryFind_2(arrary *[100]int, leftIndex int, rightIndex int, findVal int) { //查询函数
	middle := (leftIndex + rightIndex) / 2

	if leftIndex > rightIndex {
		fmt.Printf("没有找到")
		return
	}

	if (*arrary)[middle] > findVal {
		BinaryFind_2(arrary, leftIndex, middle-1, findVal)
	} else {
		fmt.Println("找到了是：", arrary[middle])
		var caRd Atm
		fmt.Println("现在的余额为：", caRd.card_3)
	}
}

func BinaryFind_3(arrary *[100]int, leftIndex int, rightIndex int, findVal int, money_2 int) { //转账函数
	middle := (leftIndex + rightIndex) / 2

	if leftIndex > rightIndex {
		fmt.Printf("没有找到")
		return
	}

	if (*arrary)[middle] > findVal {
		BinaryFind(arrary, leftIndex, middle-1, findVal, money_2)
	} else {
		fmt.Println("找到了是：", arrary[middle])
		var caRd Atm
		//var temp int64
		caRd.card_5 = money_2 + caRd.card_5
		fmt.Println("转账成功，现在本用户余额为：", caRd.card_3-money_2)
		fmt.Println("转账成功，现在被转账用户余额为：", caRd.card_5)
	}
}

//下面为添加数据，由于时间关系，本人只是将台南佳数据的代码放在下面
//type Atm struct{
//	card_1 map[string]string//map中包括卡号，金额
//	card_4 int//存款
//	card_2 int//取款
//	card_3 int//余额
//	card_5 int//转账
//}
//var card Atm
//card.card_1 = make(map[string]string)
//card.card_1["刘备"] = "123456"
//card.card_1["孙权"] = "123457"
//card.card_1["曹操"] = "123458"
//card.card_1["司马懿"] = "123459"
//card.card_1["张春华"] = "123466"
//card.card_1["张星彩"] = "123467"
//card.card_1["张让"] = "123468"
//card.card_1["刘赞"] = "123469"
//card.card_1["诸葛亮"] = "123470"
//card.card_1["黄月英"] = "123471"
