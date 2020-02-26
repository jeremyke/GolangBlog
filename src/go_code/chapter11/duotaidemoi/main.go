package main

import ()

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	//b = a//这样直接赋值是不对的，因为a是空接口类型，而b是point类型，这时需要类型断言
	b = a.(Point) //类型断言
}
