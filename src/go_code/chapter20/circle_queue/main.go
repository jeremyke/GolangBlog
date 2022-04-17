package main

import "fmt"

//数组模拟单向队列
type circleQueue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

func (aq *circleQueue) addQueue(num int) bool {
	if (aq.rear+1)%aq.maxSize == aq.front {
		return false
	}
	aq.array[aq.rear] = num
	aq.rear = (aq.rear + 1) % aq.maxSize
	return true
}

func (aq *circleQueue) deleteQueue() int {
	if aq.front == aq.rear {
		return -1
	}
	res := aq.array[aq.front]
	aq.front = (aq.front + 1) % aq.maxSize
	return res
}

func (aq *circleQueue) showQueue() {
	fmt.Println(aq)
	//队列的有效数据个数
	validNum := (aq.rear + aq.maxSize - aq.front) % aq.maxSize
	for i := aq.front; i < aq.front+validNum; i++ {
		fmt.Printf("%d ", aq.array[i%aq.maxSize])
	}
	fmt.Println()
}

func main() {
	var myQueue = &circleQueue{
		maxSize: 5,
		front:   0,
		rear:    0,
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
