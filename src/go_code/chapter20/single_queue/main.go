package main

import "fmt"

//数组模拟单向队列
type arrQueue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

func (aq *arrQueue) addQueue(num int) bool {
	if aq.rear == aq.maxSize-1 {
		return false
	}
	aq.rear++
	aq.array[aq.rear] = num
	return true
}

func (aq *arrQueue) deleteQueue() int {
	if aq.front == aq.rear {
		return -1
	}
	aq.front++
	res := aq.array[aq.front]
	return res
}

func (aq *arrQueue) showQueue() {
	for i := aq.front + 1; i <= aq.rear; i++ {
		fmt.Printf("%d ", aq.array[i])
	}
	fmt.Println()
}

func main() {
	var myQueue = arrQueue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	for {
		var selectOpt int
		fmt.Println("-------请输入你的选项------")
		fmt.Println("1.添加数据")
		fmt.Println("2.取出数据")
		fmt.Println("3.展示队列")
		fmt.Scanf("%d\n", &selectOpt)
		switch selectOpt {
		case 1:
			var addNum int
			fmt.Println("请输入添加的数据")
			fmt.Scanf("%d\n", &addNum)
			addRes := myQueue.addQueue(addNum)
			if addRes {
				fmt.Println("添加成功")
			} else {
				fmt.Println("添加失败")
			}
		case 2:
			resNum := myQueue.deleteQueue()
			fmt.Println("取出的数据是:", resNum)
		case 3:
			myQueue.showQueue()
		default:
			fmt.Println("输入错误")
		}
	}

}
