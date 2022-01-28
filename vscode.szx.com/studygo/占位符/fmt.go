package main

import "fmt"

//fmt占位符
func main() {
	var n = 100
	//查看类型
	fmt.Printf("%T\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%v\n", n)

	m := "hello,Go语言！"
	fmt.Printf("%s\n", m)
	fmt.Printf("字符串：%v\n", m)
	fmt.Printf("%T value:%v\n", m, m)
	fmt.Printf("%#v", m)
}

/*
	Printf("格式化字符串",值)
	%T：查看类型
	%d：十进制数
	%b：二进制数
	%o：八进制数
	%x：十六进制数
	%c：字符
	%s：字符串
	%p：指针
	%v：值
	%f：浮点数
*/
