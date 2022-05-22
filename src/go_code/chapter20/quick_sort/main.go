package main

import "fmt"

func quickSort(left, right int, paramArr *[]int) {
	l := left
	r := right
	//找到中间点
	pivot := (*paramArr)[(left+right)/2]
	//for循环把小于pivot的数放在左边；大的放在右边
	for l < r {
		//先从pivot左边找到大于等于pivot的数
		for (*paramArr)[l] < pivot {
			l++
		}
		//先从pivot右边找到小于等于pivot的数
		for (*paramArr)[r] > pivot {
			r--
		}
		//本次分解任务完成
		if l >= r {
			break
		}
		//找到左边大于中间值的数和右边小于中间值的数据后， 开始交换
		(*paramArr)[l], (*paramArr)[r] = (*paramArr)[r], (*paramArr)[l]

		//交换后，如果左边的值等于中间值的数，下次就不比较
		if (*paramArr)[l] == pivot {
			r--
		}
		//交换后，如果右边的值等于中间值的数，下次就不比较
		if (*paramArr)[r] == pivot {
			l++
		}
	}
	//如果l==r再移动一位
	if l == r {
		l++
		r--
	}
	//递归
	if left < r {
		quickSort(left, r, paramArr)
	}
	if right > l {
		quickSort(l, right, paramArr)
	}
}

func main() {
	aa := []int{7, 3, 8, 1, 9, 4, 10}
	quickSort(0, len(aa)-1, &aa)
	fmt.Println(aa)
}
