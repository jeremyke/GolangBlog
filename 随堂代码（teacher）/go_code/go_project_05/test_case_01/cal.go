package main
import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T){
	res := addUpper(10)
	if res != 55{
		t.Fatalf("错误", 55, res)
	}
	t.Logf("正确")
}