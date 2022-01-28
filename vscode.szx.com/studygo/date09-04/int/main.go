package main

import "fmt"

//整型——进制
//go语言中没办法直接定义二进制数

func main() {
	//十进制
	var i1 = 101
	fmt.Printf("i1:"+"%d\n", i1)
	fmt.Printf("%b\n", i1) //把十进制转化成二进制
	fmt.Printf("%o\n", i1) //把十进制转换成八进制
	fmt.Printf("%x\n", i1) //把十进制转换成十六进制

	//八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	//十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	//查看变量类型
	fmt.Printf("%T\n", i3)

	//声明int8类型的变量
	i4 := int8(9) //明确指定int8类型，否则就是默认int类型
	fmt.Printf("i4的类型："+"%T\n", i4)
	fmt.Println(i4)
}
