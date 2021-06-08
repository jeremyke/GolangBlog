package main

import "fmt"

func main() {
	init_arr := [10]byte{'a', 'c', 'd', 'e', 'f', 'g', 'h', 'k', 'o'}
	erfen(init_arr, 'e', 0, 9)
}

func erfen(arr [10]byte, option byte, leftIndex int, rightIndex int) {
	if leftIndex > rightIndex {
		fmt.Println("不存在找不到")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	fmt.Printf("arr[middle]=%v\n", arr[middle])
	fmt.Printf("option=%v\n", option)
	if arr[middle] > option {
		erfen(arr, option, leftIndex, middle-1)
	} else if arr[middle] < option {
		erfen(arr, option, middle+1, rightIndex)
	} else {
		fmt.Println("找到了，下标为：", middle)
	}
}
