package service

import (
	"go_code/customerManager/model"
)

//对customer的操作
type CustomerService struct {
	customers   []model.Customer
	customerNum int //客户数
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张山", "男", 25, "15013526701", "zs@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

func (cs *CustomerService) Add(customer model.Customer) bool {
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

func (cs *CustomerService) Delete(id int) bool {
	cs.customerNum--
	cs.customers = append(cs.customers[:id], cs.customers[id+1:]...)
	return true
}

func (cs *CustomerService) Edit(k int, customer model.Customer) bool {
	cs.customers[k] = customer
	return true
}
