package model

type Student struct {
	Name  string
	Score float64
}

type teacher struct {
	name    string
	Subject string
}

func NewTeacher(name, subject string) *teacher {
	return &teacher{
		name:    name,
		Subject: subject,
	}
}

//如果Name首写字母小写，则在其他包不可以直接使用，我们可以提供一个方法

func (t *teacher) GetName() string {
	return t.name
}
