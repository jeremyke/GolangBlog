package main

import (
	"GolangBlog/src/go_code/chapter20/hashtable/emphashtable"
	"GolangBlog/src/go_code/chapter20/hashtable/employ"
	"fmt"
	"os"
)

func main() {
	key := ""
	id := 0
	name := ""
	hashtable := new(emphashtable.EmpHashTable)
	for {
		fmt.Println("===============雇员系统菜单============")
		fmt.Println("input 添加")
		fmt.Println("show  显示")
		fmt.Println("find  查找")
		fmt.Println("exit  退出")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			Add(hashtable, id, name)
		case "show":
			Show(hashtable)
		case "find":
			fmt.Println("请输入id号:")
			fmt.Scanln(&id)
			Find(hashtable, id)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}
}

func Add(hashtable *emphashtable.EmpHashTable, id int, name string) {
	emp := &employ.Emp{
		Id:   id,
		Name: name,
	}
	hashtable.Insert(emp)
}

func Show(hashtable *emphashtable.EmpHashTable) {
	hashtable.ShowAll()
}

func Find(hashtable *emphashtable.EmpHashTable, id int) {
	emp := hashtable.FindById(id)
	if emp == nil {
		fmt.Printf("id=%d 的雇员不存在\n", id)
	} else {
		//编写一个方法，显示雇员信息
		emp.ShowMe()
	}
}
