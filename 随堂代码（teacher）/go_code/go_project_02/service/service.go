package service
import (
	"fmt"
	"go_code\go_project_02\model"
)

type CustomerService struct{
	customers []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService{
	//为了能看到有用户再切片中，先初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(...)
	customerService.custmoers = append(customerService.customers,customer)
	return customerService
}	
//返回用户切片
func (this *CustomerService) List() []model.Customer{
	return this.costomers
}
func (this *CustomerService) Add(customer mdoel.Customer){
	//我们确定一个规则，确定Id的顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.custmoers = append(this.custmoers,custmoer)
	return true
}
