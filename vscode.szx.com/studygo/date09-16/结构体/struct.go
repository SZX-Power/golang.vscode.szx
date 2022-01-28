package main

import "fmt"

//结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	//声明一个person类型的变量p
	var p person
	// 通过字段赋值
	p.name = "魏无羡"
	p.age = 18
	p.gender = "男"
	p.hobby = []string{"吹笛", "饮酒", "舞剑"}
	fmt.Println(p)
	//访问变量p的字段
	fmt.Printf("%T\n", p)
	fmt.Println(p.name)
	var p2 person
	p2.name = "szx"
	p2.age = 18
	fmt.Printf("%T %v\n", p2, p2)

	//匿名结构体：多用于临时场景
	var s struct {
		x string
		y int
	}
	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("%T %v\n", s, s)
}
