package main

import "fmt"

/*
     //Go语言函数外的语句必须以关键字开头
	 var name = "娜扎"
	 const age = 18

	 //age = 18  无关键字不可以
*/

// main函数是入口函数，他没有参数也没有返回值
func main() {
	//函数内部定义、声明的变量必须使用
	var isOK = true
	const a = 2
	fmt.Println(isOK)

	var name string
	name = "史"

	name1 := "史振兴" //只能在函数内部使用此法
	fmt.Println(name, name1)
}

/*
	整型：
		无符号整型：uint8、uint16、uint32、uint64
		带符号整型：int8、int32、int64
		uint和int：具体是32位还是64位看操作系统
		uintptr：表示指针
	浮点型：
		float64和float32
		Go语言中浮点数默认是float64
	复数：
		complex128和complex64
	布尔值：
		true和false
		不能和其他的类型做转换
	字符串：
		常用方法
		字符串不能修改
	进制：
		go语言中没办法直接定义二进制数
		可以定义八进制（0）和十六进制（0x）十进制
	byte和rune类型：
		都属于类型别名
	字符串、字符、字节都是什么？
		字符串：双引号包裹的是字符串
		字符：单引号包裹的是字符、如单个字母、单个符号、单个汉字
		字节：1byte=8bit
	go语言中字符串都是UTF8编码，UTF8编码中一个常用汉字一般占用3个字节。
*/
