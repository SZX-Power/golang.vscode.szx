package main

import "fmt"

//defer
//defer多用于函数结束之前释放资源（文件句柄、数据库连接、socket连接）

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿") // defer把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("呵呵呵") // 一个函数可以有多个defer语句
	defer fmt.Println("哈哈哈") // 多个defer语句按照（后进先出、先进后出）的顺序延迟执行;将延迟处理的语句按defer定义的逆序进行执行
	fmt.Println("end")
}

func main() {
	deferDemo()
}
