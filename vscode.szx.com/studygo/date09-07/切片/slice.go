package main

import "fmt"

//切片Slice
func main() {
	//切片的定义
	var s1 []int    // 定义一个存放int类型元素的切片
	var s2 []string // 定义一个存放string类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // true   nil（空，相当于没有开辟内存空间）
	fmt.Println(s2 == nil) // true
	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "长江", "黄海"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // false   nil（空，相当于没有开辟内存空间）
	fmt.Println(s2 == nil) // false
	//长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	//由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	a2 := a1[0:4] //基于一个数组切割，左包含右不包含，左闭右开
	fmt.Println(a2)
	a3 := a1[1:6]
	fmt.Println(a3)
	a4 := a1[:4]
	a5 := a1[3:] //[7 9 11 13]
	a6 := a1[:]
	fmt.Println(a4, a5, a6)

	/*
		切片指向了一个底层的数组。
		切片的长度就是它元素的个数。
		切片的容量是底层数组从切片的第一个元素到最后一个元素的数量。
	*/
	fmt.Printf("len(a4):%d cap(a4):%d\n", len(a4), cap(a4))
	fmt.Printf("len(a5):%d cap(a5):%d\n", len(a5), cap(a5))

	//切片再切割
	a7 := a5[3:] //[13]
	fmt.Printf("len(a7):%d cap(a7):%d\n", len(a7), cap(a7))
	//切片是引用类型，都指向了底层的一个数组。
	fmt.Println("a5:", a5)
	a1[6] = 1300 //修改了底层数组的值
	fmt.Println("a5:", a5)
	fmt.Println("a7:", a7)
}
