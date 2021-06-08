package main

//type A struct {
//	Name string
//	Age int
//}
//
//func (a A) say()  {
//	fmt.Println("我是属于A的")
//}
//
//type B struct {
//	Name string
//	A
//}
//
//func (b B) say() {
//	fmt.Println("我是属于B的")
//}
//
//func main() {
//	var b B
//	b.say()//我是属于B的
//	b.A.say()//我是属于A的
//}

type A struct {
	Name string
	Age  int
}
type B struct {
	Name  string
	Score int
}
type C struct {
	A
	B
}

func main() {
	var c_struct C
	c_struct.A.Name = "sss"
}
