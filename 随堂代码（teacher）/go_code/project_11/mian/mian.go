package mian

import (
	"fmt"
)

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}
func main() {
	res := fbn(9)
	fmt.Println("斐波那楔数列：", res)
}
