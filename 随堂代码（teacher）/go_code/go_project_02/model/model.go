package model

import (
	"fmt"
)

type Customer struct {
	Id     int
	Name   string
	Grnder string
	Age    int
	Phone  string
	Email  string
}

//工厂模式,返回custmer的模式
func NewCustomer(id int, name string,
	age int, phone string,
	email string) Customer {
	return Customer{
		Id:    id,
		Name:  name,
		Age:   age,
		Phone: phone,
		Email: email,
	}
}

func NewCustomer_2(name string,
	age int, phone string,
	email string) Customer {
	return Customer{
		//Id : id,
		Name:  name,
		Age:   age,
		Phone: phone,
		Email: email,
	}
}

func (this *Customer) GetInfor() string {
	info := fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name,
		this, Grnder, this.Phone, this.Email)
	return info
}
