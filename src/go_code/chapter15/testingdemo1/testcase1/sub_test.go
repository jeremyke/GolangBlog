package main

import (
	"testing"
)

func TestGetSub(t *testing.T) {
	res := getSub(10, 2)
	if res != 8 {
		t.Fatalf("getSub(10,2)执行错误，期望值=%v 实际值=%v\n", 8, res)
	}
	//正确就输出日志
	t.Logf("getSub(10,2)执行正确")
}
