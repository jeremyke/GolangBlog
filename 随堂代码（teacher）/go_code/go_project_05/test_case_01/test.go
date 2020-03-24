package test
import (
	"fmt"
)

func addUpper(n int) int{
	res := 0
	for i := 1; i <= n - 1; i++{
		res += 1
	}
	return res
}