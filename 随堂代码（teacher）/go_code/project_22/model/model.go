package model

//import "fmt"

type student struct {
	Name  string
	Score float64
}

type person struct {
	Name string  //其他包可以访问
	age  int     //其他包不可以访问
	sal  float64 //其他包不可以访问
}

//工厂模式函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了访问age 和 sal 我们编写一对SetXxxx和GetXxxx方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} // else {
	//fmt.Println("年龄不匹配")
	//}
}
func (p *person) GetAge() int {
	return p.age
}
func (p *person) SetSal(sal float64) {
	if sal > 30000 && sal < 150000 {
		p.sal = sal
	} // else {
	//	fmt.Println("薪水不匹配")
	//}
}
func (p *person) GetSal() float64 {
	return p.sal
}

func NewStudent(name string, score float64) *student {
	return &student{
		Name:  name,
		Score: score,
	}
	//fmt.Println(name,  score)
}
