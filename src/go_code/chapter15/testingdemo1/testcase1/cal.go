package main

import ()

func AddUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func getSub(n, m int) int {
	return n - m
}
