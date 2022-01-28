package main

import "fmt"

var name string

var na string = "szx"

//go是强语言类型 定义数组必须加类型。
//go是强类型语言，但任何类型都可以转为interfalce{}类型。 所以用 []interface{}
type szx interface {
	//go如何实现这种存放不同类型的数组
}

func main() {
	number := 101001
	//name = "sz"
	b := [...]szx{1, 2, 3.14, "hello", true}
	fmt.Println(b)
	fmt.Println(name, na)
	fmt.Println(number)
}
