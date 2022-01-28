package main

import "fmt"

//自定义类型和类型别名
//type后面跟的是类型

type myInt int //自定义类型
type your int  // 类型别名

func main() {
	var n myInt
	n = 7
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	var m your
	m = 100
	fmt.Printf("%d %T\n", m, m)

	var c rune
	c = '中'
	fmt.Printf("%d %T\n", c, c)
}
