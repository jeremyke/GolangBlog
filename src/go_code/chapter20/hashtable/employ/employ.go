package employ

import "fmt"

//Emp 员工结构体
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (e *Emp) ShowMe() {
	fmt.Printf("链表%d 找到该雇员 %d\n", e.Id%7, e.Id)
}
