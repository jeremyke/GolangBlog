package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s *Student) showInfo() {
	fmt.Printf("学生名=%v,年龄=%v，成绩=%v", s.Name, s.Age, s.Score)
}
func (s *Student) SetScore(score int) {
	s.Score = score
}

type smallStu struct {
	Student
}

func (s *smallStu) test() {
	fmt.Println("小学生在考试...")
}

type bigStu struct {
	Student
}

func (b *bigStu) test() {
	fmt.Println("大学生在考试...")
}

func main() {
	smallStu1 := &smallStu{}
	smallStu1.Student.Name = "tom"
	smallStu1.Age = 7
	smallStu1.test()
	smallStu1.Student.SetScore(89)
	smallStu1.Student.showInfo()
}
