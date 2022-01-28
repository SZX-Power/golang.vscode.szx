package main

import (
	"fmt"
	"sort"
)

// 切片的练习题
func main() {
	var a = make([]int, 5, 10) // 创建切片，长度为5，容量为10
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	fmt.Println(cap(a))

	// sort对切片进行排序
	var a1 = [...]int{1, 3, 8, 2, 9, 1}
	sort.Ints(a1[:])
	fmt.Println(a1)

}
