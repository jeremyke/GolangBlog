package view
import (
	"fmt"
	"go_code\go_project_02\service"
)

type customerView struct{
	key string
	loop bool//表示是否循环线是主菜单
	customerService *CustomerService
}

func (this *customerView) list(){
	this.customers := this.customerService.List()
	//显示
	fmt.Println("-------------客户列表------------")
	fmt.Println("编号  姓名  性别  年龄  电话  邮箱")
	for i := 1; i <= len(customer); i++{
		fmt.Println(customer[i].GetInfo())
	}
	fmt.Println("-----------客户列表完成----------")
	//fmt.Println("-------------客户列表------------")

}
//得到用户的输入信息构建新的用户
func (this *customerView) add(){
	fmt.Println("-------------添加客户-------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	...

	customer :=model.NewCustomer_2(name...)
	this.customerService.Add(customer)
}

func (cv *customerView) mainMenu(){
	for {
		fmt.Println("-------------客户信息-------------")
		fmt.Println("-------------添加客户-------------")
		fmt.Println("-------------修改客户-------------")
		fmt.Println("-------------删除客户-------------")
		fmt.Println("-------------客户列表-------------")
		fmt.Println("---------------退出---------------")
		fmt.Println("-----------请选择（1-5）：-----------")

		fmt.Scanln(&this.key)
		switch this.key {
		case 1:
			//fmt.Println("添加客户")
			this.add()
		case 2:
			fmt.Println("修改客户")
		case 3:
			fmt.Println("删除客户")
		case 4:
			//fmt.Println("客户列表")
			this.list()
		case 5:
			this.loop = false
			//fmt.Println("退出")
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
		if this.loop = false {
			break
		}
		
	}

	fmt.Println("感谢使用")
}

func main(){
	customerView := customerView{
		key : "",
		loop : true,
	}
	//完成对customerView的初始化
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
