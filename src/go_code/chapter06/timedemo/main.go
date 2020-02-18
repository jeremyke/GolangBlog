package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("now=%v type=%T\n", now, now)
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf(now.Format("2006/01/02 15:04:05\n"))
	fmt.Printf(now.Format("01"))
}
