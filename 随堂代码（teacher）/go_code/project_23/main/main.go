package main

import (
	"fmt"
)

type Pupil struct {
	Name  string
	Age   int
	Score int
}

func (p *Pupil) ShowInfo() {
	fmt.Println("测试")
}

func (p *Pupil) SetScore(score int) {
	//此处应有业务判断
	p.Score = score
}

func (p *Pupil) testing() {
	fmt.Println("正在考试中")
}

func main() {
	var pupil = &Pupil{
		Name:  "伊莉雅",
		Age:   10,
		Score: 96,
	}
	pupil.testing()
	pupil.SetScore(96)
	pupil.ShowInfo()
}
