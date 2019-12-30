package main

import (
	"fmt"
	"go_code/golicense/chapter2/tempconv"
)

func main() {
	fmt.Printf("Brrrr!%v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}
