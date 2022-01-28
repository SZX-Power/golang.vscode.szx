package main

import "fmt"

//匿名字段
//适用于字段比较少也比较简单的场景
//不常用！！！
type person struct {
	string
	int
}

func main() {
	p1 := person{
		"周琳",
		17,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
