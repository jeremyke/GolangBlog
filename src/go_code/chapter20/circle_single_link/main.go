package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func main() {
	headNode := &CatNode{}

	node1 := &CatNode{no: 1, name: "tom"}
	node2 := &CatNode{no: 2, name: "jerry"}
	node3 := &CatNode{no: 3, name: "cate"}
	node4 := &CatNode{no: 4, name: "poul"}

	AddCate(headNode, node1)
	AddCate(headNode, node2)
	AddCate(headNode, node3)
	AddCate(headNode, node4)
	showCateNode(headNode)
	headNode = delCateNode(headNode, node4)
	showCateNode(headNode)
}

func AddCate(head, NewCat *CatNode) {
	//判断是否是头节点
	if head.next == nil {
		head.no = NewCat.no
		head.name = NewCat.name
		head.next = head
		return
	}
	tempNode := head
	for {
		if tempNode.next == head {
			break
		}
		tempNode = tempNode.next
	}
	//加入链表
	tempNode.next = NewCat
	NewCat.next = head

}

func delCateNode(head, delNod *CatNode) (headNode *CatNode) {
	headNode = head
	if head.next == nil {
		fmt.Println("链表为空")
		return
	}

	//1个元素的链表
	if head.next == head {
		if head.no == delNod.no {
			head.no = 0
			head.name = ""
			head.next = nil
			return
		} else {
			fmt.Println("未找到要删除的节点")
			return
		}
	}

	//多个元素的链表
	tempNode := head
	rearNode := head
	for {
		if rearNode.next == head {
			break
		}
		rearNode = rearNode.next
	}

	for {
		//头节点就是需要删除的节点
		if head.no == delNod.no {
			headNode = head.next
			rearNode.next = headNode
			break
		}
		//中间节点是需要删除的节点
		if tempNode.next == delNod {
			tempNode.next = delNod.next
			break
		}
		//尾节点是需要删除的节点
		if tempNode.next == head {
			if tempNode.no == delNod.no {
				//找到倒数第二个节点
				descSecondNode := head
				for {
					if descSecondNode.next == tempNode {
						break
					}
					descSecondNode = descSecondNode.next
				}
				descSecondNode.next = head
			} else {
				fmt.Println("未找到要删除的节点")
				return
			}
			break
		}
		tempNode = tempNode.next
	}
	return
}

func showCateNode(head *CatNode) {
	tempNode := head
	if tempNode.next == nil {
		fmt.Println("链表为空")
		return
	}
	for {
		fmt.Printf("%d %s-->", tempNode.no, tempNode.name)
		if tempNode.next == head {
			break
		}
		tempNode = tempNode.next
	}
	fmt.Println()
}
