package main

import "fmt"

type Goods struct {
	Name  string
	Price float64
}

type Book struct {
	Authod string
	Goods
}

//继承了2个stuct
type MathBook struct {
	Book
	CaculateNum int
}

func main() {
	var mathBook MathBook
	mathBook.CaculateNum = 11
	mathBook.Authod = "祖冲之"
	mathBook.Name = "九章算术"
	mathBook.Price = 50.00
	fmt.Println(mathBook)
}
