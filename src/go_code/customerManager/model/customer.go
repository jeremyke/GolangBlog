package model

//定义一个结构体 表示客户信息

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//编写工厂方法初始化结构体

func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func NewCustomer2(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
