package employlink

import (
	"GolangBlog/src/go_code/chapter20/hashtable/employ"
	"fmt"
)

//EmpLink 员工链表
type EmpLink struct {
	Head *employ.Emp
}

//Insert 添加员工的方法, 保证添加时，编号从小到大
func (el *EmpLink) Insert(emp *employ.Emp) {
	cur := el.Head            // 这是辅助指针
	var pre *employ.Emp = nil // 这是一个辅助指针 pre 在cur前面

	//如果当前的EmpLink就是一个空链表
	if cur == nil {
		el.Head = emp //完成
		return
	}
	//如果不是一个空链表,给emp找到对应的位置并插入：思路是 让 cur 和 emp 比较，然后让pre 保持在 cur 前面
	//处理头结点
	if cur.Id > emp.Id {
		oldHead := el.Head
		emp.Next = oldHead
		el.Head = emp
		return
	}
	//处理后续节点
	for {
		if cur != nil {
			if cur.Id > emp.Id { //当前位置id大于即将插入的id
				//找到位置
				break
			} else if cur.Id == emp.Id { //当前位置id等于即将插入的id
				fmt.Printf("id为%d的员工已存在，请重新输入id\n", emp.Id)
				return
			}
			pre = cur //保证同步
			cur = cur.Next
		} else {
			break
		}
	}
	//退出时，我们看下是否将emp添加到链表最后
	pre.Next = emp
	emp.Next = cur

}

//ShowLink 显示链表的信息
func (el *EmpLink) ShowLink(no int) {
	if el.Head == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}
	//变量当前的链表，并显示数据
	cur := el.Head // 辅助的指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println() //换行处理
}

//FindById 根据id查找对应的雇员，如果没有就返回nil
func (el *EmpLink) FindById(id int) *employ.Emp {
	cur := el.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}
