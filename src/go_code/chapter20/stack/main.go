package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int    //栈最大可以存放的数的个数
	Top    int    //表示栈顶，因为栈顶固定，因此我们可以直接使用Top
	arr    [5]int //数组模拟栈
}

func main() {
	myStack := Stack{
		MaxTop: 5,
		Top:    -1, //默认值，栈为空
	}
	myStack.Push(1)
	myStack.Push(4)
	myStack.Push(9)
	fmt.Println("top是：", myStack.Top)
	myStack.List()

	fmt.Println("开始pop1个")
	val, _ := myStack.Pop()
	fmt.Println("top是：", myStack.Top)
	fmt.Println("pop值：", val)
	myStack.List()

	fmt.Println("再pop1个")
	val2, _ := myStack.Pop()
	fmt.Println("top是：", myStack.Top)
	fmt.Println("pop值：", val2)
	myStack.List()

	fmt.Println("再pop1个")
	val3, _ := myStack.Pop()
	fmt.Println("top是：", myStack.Top)
	fmt.Println("pop值：", val3)
	myStack.List()

}

func (s *Stack) Push(val int) error {
	if s.Top == s.MaxTop-1 {
		return errors.New("栈满了")
	}

	s.Top++
	s.arr[s.Top] = val

	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.Top == -1 {
		return 0, errors.New("空栈")
	}

	res := s.arr[s.Top]
	s.arr[s.Top] = 0
	s.Top--
	return res, nil
}

func (s *Stack) List() {
	if s.Top == -1 {
		fmt.Println("空栈")
	}
	for i := 0; i < s.Top+1; i++ {
		fmt.Printf("arr[%v]=%v\n", i, s.arr[i])
	}
}
