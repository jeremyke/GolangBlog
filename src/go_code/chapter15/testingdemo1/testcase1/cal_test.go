package main

import (
	_ "fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10)执行错误，期望值=%v 实际值=%v\n",55,res)
		t.Fatalf("AddUpper(10)执行错误，期望值=%v 实际值=%v\n", 55, res)
	}
	//正确就输出日志
	t.Logf("AddUpper(10)执行正确")
}
