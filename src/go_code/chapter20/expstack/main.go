package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

type Stack struct {
	MaxTop int     //栈最大可以存放的数的个数
	Top    int     //表示栈顶
	arr    [20]int //数组模拟栈
}

var itobMap = map[int]byte{
	1: '+',
	2: '-',
	3: '*',
	4: '/',
}
var btoiMap = map[byte]int{
	'+': 1,
	'-': 2,
	'*': 3,
	'/': 4,
}

var opePriorityMap = map[byte]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
}

func main() {
	expStr := "3+2*6-2"
	expStrArr := []byte(expStr)
	a := operation(expStrArr)
	fmt.Println(a)

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

func operation(expStr []byte) int {
	numStack := Stack{
		MaxTop: 20,
		Top:    -1,
	}
	opeStack := Stack{
		MaxTop: 20,
		Top:    -1,
	}
	//入栈
	for _, v := range expStr {
		vInt, _ := strconv.Atoi(string(v))
		if unicode.IsDigit(rune(v)) {
			numStack.Push(vInt)
		} else {
			if opeStack.Top == -1 {
				opeStack.Push(btoiMap[v])
			} else {
				if opeStack.arr[opeStack.Top] >= opePriorityMap[v] {
					opeVal, _ := opeStack.Pop()
					numVal2, _ := numStack.Pop()
					numVal1, _ := numStack.Pop()
					resTmp := simpleOpe(numVal1, numVal2, itobMap[opeVal])
					numStack.Push(resTmp)
					opeStack.Push(btoiMap[v])
				} else {
					opeStack.Push(btoiMap[v])
				}
			}
		}
	}
	//符号出栈
	for opeStack.Top >= 0 {
		opeTmp, _ := opeStack.Pop()
		numVal2, _ := numStack.Pop()
		numVal1, _ := numStack.Pop()
		resTmp := simpleOpe(numVal1, numVal2, itobMap[opeTmp])
		numStack.Push(resTmp)
	}

	res, _ := numStack.Pop()
	return res
}

func simpleOpe(a, b int, operation byte) int {
	if operation == '+' {
		return a + b
	} else if operation == '-' {
		return a - b
	} else if operation == '*' {
		return a * b
	} else {
		return a / b
	}
}
