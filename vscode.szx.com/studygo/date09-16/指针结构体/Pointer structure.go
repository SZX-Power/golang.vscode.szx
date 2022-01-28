package main

import "fmt"

//结构体是值类型

type person struct {
	name, gender string
	age          int
}

// go语言中函数传参数永远传是拷贝
func f(x person) {
	x.gender = "女" // 修改的是副本的gender

}

func f1(x *person) {
	(*x).gender = "女" //根据内存地址找到那个原变量，修改的是原来的变量
	x.gender = "女"    // 语法上，自动根据指针找对应的变量
}

func main() {
	var p person
	p.name = "史振兴"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender) // 男

	f1(&p)                // 内存地址，像这种0x123ec2
	fmt.Println(p.gender) // 女

	var x *person
	fmt.Println(x)

	//1、结构体指针1
	var p2 = new(person)
	p2.name = "理想"
	(*p2).gender = "女"

	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)  // p2保存的值就是一个内存地址
	fmt.Printf("%p\n", &p2) // 求b2的内存地址

	//2、结构体指针2
	//2.1 key-value初始化
	var p3 = &person{
		name: "元帅",
	}
	fmt.Printf("%#v\n", p3)

	//2.2 使用值列表的形式初始化，值的顺序要和结构体定义时字段的顺序一致
	p4 := person{
		"小王子",
		"男",
		18,
	}
	fmt.Printf("%#v\n", p4)
	/*
		使用这种格式初始化时，需要注意：
		必须初始化结构体的所有字段。
		初始值的填充顺序必须与字段在结构体中的声明顺序一致。
		该方式不能和键值初始化方式混用。
	*/
}
