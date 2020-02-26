package person

type person struct {
	name    string
	age     int
	sallary float64
}

//类构造函数
func NewPerson(name string, age int, sallary float64) *person {
	return &person{
		name:    name,
		age:     age,
		sallary: sallary,
	}
}

func (p *person) SetAge(age int) {
	p.age = age
}

func (p *person) GetAge() int {
	return p.age
}
