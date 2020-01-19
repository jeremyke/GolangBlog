package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

//定义一个结构体切片
type StudentSlice []Student

func (s StudentSlice) Len() int {
	return len(s)
}
func (s StudentSlice) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}
func (s StudentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//对一个结构体切片的某个元素排序
func main() {
	//1.定义一个切片,且赋值
	var myStudent StudentSlice
	for i := 0; i < 10; i++ {
		aStudent := Student{
			Name:  fmt.Sprintf("学生--%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: rand.Float64(),
		}
		myStudent = append(myStudent, aStudent)
	}
	//2.排序前的顺序
	fmt.Println("----------------排序前-----------------------")
	for _, v := range myStudent {
		fmt.Println(v)
	}
	//3.排序
	sort.Sort(myStudent)
	//4.排序后的顺序
	fmt.Println("----------------排序后-----------------------")
	for _, v := range myStudent {
		fmt.Println(v)
	}
}
