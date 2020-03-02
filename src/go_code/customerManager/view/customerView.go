package main

import (
	"fmt"
	"go_code/customerManager/model"
	"go_code/customerManager/service"
)

type customerView struct {
	key             string //选择序号
	loop            bool
	customerService *service.CustomerService
}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.MainMenu()
}

//显示主菜单

func (cv *customerView) MainMenu() {
	for {
		fmt.Println("-------客户信息管理系统-------")
		fmt.Println("          1 添加客户")
		fmt.Println("          2 修改客户")
		fmt.Println("          3 删除客户")
		fmt.Println("          4 客户列表")
		fmt.Println("          5 退    出")
		fmt.Println("请选择（1-5）：")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.edit()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
		if !cv.loop {
			fmt.Println("退出成功...")
			break
		}
	}
}

//客户列表
func (cv *customerView) list() {
	customers := cv.customerService.List()
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮件")
	for _, v := range customers {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", v.Id, v.Name, v.Gender, v.Age, v.Phone, v.Email)
	}
	fmt.Println("-------------------------客户列表完成---------------------------")
}

//添加客户
func (cv *customerView) add() {
	fmt.Println("---------------------------添加用户---------------------------")
	var name string
	var gender string
	var age int
	var phone string
	var email string
	fmt.Println("请一次输入用户的姓名，性别，年龄，电话，邮件")
	fmt.Scanf("%v %v %d %v %v", &name, &gender, &age, &phone, &email)
	customers := model.NewCustomer2(name, gender, age, phone, email)
	if cv.customerService.Add(customers) {
		fmt.Println("---------------------------添加成功---------------------------")
	} else {
		fmt.Println("---------------------------添加失败---------------------------")
	}
}

//删除客户
func (cv *customerView) delete() {
	fmt.Println("---------------------------删除客户---------------------------")
	fmt.Println("请输入客户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除请输入（Y/N）：")
	choise := ""
	fmt.Scanln(&choise)
	if choise != "Y" {
		return
	}
	flag := false
	customers := cv.customerService.List()
	var k int
	var v model.Customer
	for k, v = range customers {
		if v.Id == id {
			flag = true
			break
		}
	}
	if flag {
		if cv.customerService.Delete(k) {
			fmt.Println("删除成功")
		} else {
			fmt.Println("删除失败")
		}
	} else {
		fmt.Println("---------------------------你输入的编号有误！！---------------------------")
	}
}

//修改客户
func (cv *customerView) edit() {
	fmt.Println("---------------------------修改用户---------------------------")
	fmt.Println("请输入客户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	flag := false
	customers := cv.customerService.List()
	var k int
	var v model.Customer
	for k, v = range customers {
		if v.Id == id {
			flag = true
			break
		}
	}
	if flag {
		var name string
		var gender string
		var age int
		var phone string
		var email string
		fmt.Println("请一次输入用户的姓名，性别，年龄，电话，邮件")
		fmt.Scanf("%v %v %d %v %v", &name, &gender, &age, &phone, &email)
		customers := model.NewCustomer2(name, gender, age, phone, email)
		if cv.customerService.Edit(k, customers) {
			fmt.Println("---------------------------修改成功---------------------------")
		} else {
			fmt.Println("---------------------------修改失败---------------------------")
		}
	} else {
		fmt.Println("---------------------------你输入的编号有误！！---------------------------")
	}
}
