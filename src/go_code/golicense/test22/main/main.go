package main

import (
	"fmt"
	"go_code/golicense/test22/longconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "long:%v\n", err)
		}
		m := longconv.Mile(t)
		f := longconv.Foot(t)
		fmt.Printf("%s = %s, %s = %s\n", t, longconv.MToF(m), t, longconv.FToM(f))
	}
}
