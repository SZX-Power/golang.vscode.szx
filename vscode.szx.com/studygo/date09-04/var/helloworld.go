package main

import "fmt"

//Go语言中推荐使用驼峰式命名
/*
var student_name string
var studentName string
var StudentName string
*/

//声明变量
// var name string
// var age int
// var isOK bool

//批量声明
var (
	name string // ""
	age  int    // 0
	isOK bool   //false
)

func main() {
	name = "理想"
	age = 16
	isOK = true
	//go语言中非全局变量声明必须使用，不使用就编译不过去
	fmt.Print(isOK) //在终端中输出要打印的内容
	fmt.Println()
	fmt.Printf("name:%s\n", name) //%s：占位符  使用name这个变量的值去替换占位符  \n换行
	fmt.Println(age)
	fmt.Println("Hello world!")

	//声明变量同时赋值
	var s1 string = "szx"
	fmt.Println(s1)
	//类型推导（根据值判断该变量是什么类型）
	var s2 = "23"
	fmt.Println(s2)
	//简短变量声明，只能在函数里面使用
	s3 := "170"
	fmt.Print(s3)
	//s1 :="10" //同一个作用域（{}）中不能重复声明同名的变量

	//匿名变量(哑元变量)是一个特殊的变量：_
	//当有些数据必须用变量接收但是又不使用它时，就可以用_来接收这个值。

}
