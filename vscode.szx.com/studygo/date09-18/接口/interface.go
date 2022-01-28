package main

import "fmt"

/*
	Go语言中接口（interface）是一种类型，一种抽象的类型。
	interface是一组method的集合，是duck-type programming的一种体现。接口做的事情就像是定义一个协议（规则）
*/
type cat struct{}

type dog struct{}

//定义了一个能叫的类型
type speaker interface {
	speak() // 只要实现了speak方法的变量都是speak类型
}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

func da(x speaker) {
	//接收一个参数，传进来什么，我就打什么
	x.speak() //挨打了就要叫
}
func main() {
	var c1 cat
	var d1 dog

	da(c1)
	da(d1)

	var ss speaker //定义了一个接口类型：speaker的变量：ss
	ss = c1
	fmt.Printf("%T %v\n", ss, ss)

}
