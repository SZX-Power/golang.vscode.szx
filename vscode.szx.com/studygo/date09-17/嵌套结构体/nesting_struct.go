package main

import "fmt"

// 匿名字段  嵌套结构体  匿名嵌套结构体  匿名嵌套结构体的字段冲突
// 匿名字段:结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
//嵌套结构体:一个结构体中可以嵌套包含另一个结构体或结构体指针

type address struct {
	province string
	city     string
}

type worksplace struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	addr address
	worksplace
}

type company struct {
	name    string
	address // 匿名嵌套结构体  此处类似 address:address
}

type robot struct {
	name string
	address
	worksplace
}

func main() {
	p1 := person{
		name: "史振兴",
		age:  22,
		addr: address{
			province: "山东省",
			city:     "青岛",
		},
	}
	c1 := company{
		name: "中瑞",
		address: address{
			province: "山东",
			city:     "青岛",
		},
	}
	r1 := robot{
		name: "小哪吒",
		address: address{
			province: "山东",
			city:     "青岛",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.addr.city) //嵌套结构体，一层一层的找
	fmt.Println(p1.city)
	fmt.Println(c1)
	fmt.Println(c1.city) //现在自己结构体找这个字段，找不到就去匿名嵌套的结构体中查找该字段

	fmt.Println(r1)
	//fmt.Println(r1.city) // 匿名嵌套结构体的字段冲突 worksplace与address中都有city
}
