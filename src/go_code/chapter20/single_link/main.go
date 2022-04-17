package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode
}

func AddHero(head, newHero *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			//找到了最后一个结点
			break
		}
		//如果没有找到，把当前结点的next指向临时结点
		temp = temp.next
	}
	temp.next = newHero
}

// 按照no顺序插入
func AddHeroNoAsc(head, newHero *HeroNode) {
	tempNode := head
	for {
		//找到比新结点小1的节点

		//最后一个结点
		if tempNode.next == nil {
			break
		}
		if tempNode.no < newHero.no && tempNode.next.no >= newHero.no {
			break
		}
		//如果没有找到，把当前结点的next指向临时结点
		tempNode = tempNode.next
	}
	//开始插入
	myNextNode := tempNode.next
	tempNode.next = newHero
	newHero.next = myNextNode
}

func DelNode(head, delNode *HeroNode) error {
	tempNode := head
	for {
		//找到比需要删除的节点的前一个节点
		if delNode.no == 0 {
			return fmt.Errorf("头结点不可删除")
		}
		if tempNode.next == nil {
			return fmt.Errorf("没有找到需要删除的节点")
		}
		if tempNode.next.no == delNode.no {
			break
		}
		//如果没有找到，把当前结点的next指向临时结点
		tempNode = tempNode.next
	}
	//开始删除
	tempNode.next = delNode.next
	return nil
}

func showHeroLink(head *HeroNode) {
	temp := head
	for {
		if temp == nil {
			break
		}
		fmt.Printf("[%d , %s , %s]->", temp.no, temp.name, temp.nickname)
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	head := &HeroNode{}
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero7 := &HeroNode{
		no:       7,
		name:     "AA",
		nickname: "AA!",
	}

	hero5 := &HeroNode{
		no:       5,
		name:     "BB",
		nickname: "BB!",
	}

	//增加节点
	//AddHero(head, hero1)
	//AddHero(head, hero2)
	//AddHero(head, hero3)
	AddHeroNoAsc(head, hero1)
	AddHeroNoAsc(head, hero5)
	AddHeroNoAsc(head, hero3)
	AddHeroNoAsc(head, hero2)
	AddHeroNoAsc(head, hero7)
	showHeroLink(head)
	//删除节点
	err := DelNode(head, hero3)
	if err != nil {
		fmt.Println(err.Error())
	}
	showHeroLink(head)

}
