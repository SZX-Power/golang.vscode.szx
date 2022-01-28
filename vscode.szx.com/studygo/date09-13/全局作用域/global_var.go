package main

import "fmt"

// 变量的作用域

var x = 100 //定义一个全部变量

//定义一个函数
func f1() {
	// x := 10
	name := "理想"
	/*
		函数中查找变量的顺序
		1、先在函数内部查找
		2、找不到就往函数的外面查找，一直找到全局
	*/
	fmt.Println(name, x)
}

func main() {
	f1()
}

/*
	1、全局作用域
	2、函数内作用域
	3、语句块作用域   在if、for等循环内
*/
