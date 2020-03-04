package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{
		Name: "七七",
		Age:  10,
	}
	allChan <- cat
	<-allChan
	<-allChan
	newCat := <-allChan

	fmt.Printf("newCat的类型是%T,值为%v\n", newCat, newCat)
	//类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v", a.Name)
}
