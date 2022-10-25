package emphashtable

import (
	"GolangBlog/src/go_code/chapter20/hashtable/employ"
	"GolangBlog/src/go_code/chapter20/hashtable/employlink"
)

type EmpHashTable struct {
	LinkArr [7]employlink.EmpLink
}

//Insert 给HashTable 编写Insert 雇员的方法.
func (eht *EmpHashTable) Insert(emp *employ.Emp) {
	//使用散列函数，确定将该雇员添加到哪个链表
	linkNo := eht.HashFun(emp.Id)
	//使用对应的链表添加
	eht.LinkArr[linkNo].Insert(emp)
}

//ShowAll 编写方法，显示hashtable的所有雇员
func (eht *EmpHashTable) ShowAll() {
	for i := 0; i < len(eht.LinkArr); i++ {
		eht.LinkArr[i].ShowLink(i)
	}
}

//HashFun 编写一个散列方法
func (eht *EmpHashTable) HashFun(id int) int {
	return id % 7 //得到一个值，就是对于的链表的下标
}

//FindById 编写一个方法，完成查找
func (eht *EmpHashTable) FindById(id int) *employ.Emp {
	//使用散列函数，确定将该雇员应该在哪个链表
	linkNo := eht.HashFun(id)
	return eht.LinkArr[linkNo].FindById(id)
}
