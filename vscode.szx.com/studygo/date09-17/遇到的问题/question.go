package main

import "fmt"

// 结构体遇到的问题

// 问题1、myInt(100)是个啥？
type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}
func main() {
	//声明一个int32类型的变量x，它的值是10
	/*//方法1
	var x int32 = 10 */

	/*//方法2
	var x int32
	x = 10 */

	/*//方法3
	var x = int32(10) */
	//方法4
	x := int32(10)

	fmt.Println(x)

	m := myInt(10)
	m.hello()

	//问题2、结构体初始化
	type person struct {
		name string
		age  int
	}
	//方法1：
	var p person //声明一个person类型的变量p
	p.name = "元帅"
	p.age = 18
	fmt.Println(p)
	var p1 person
	p1.name = "周琳"
	p1.age = 9000
	fmt.Println(p1)

	//方法2：键值对初始化
	p3 := person{
		name: "鸡冠花",
		age:  16,
	}
	fmt.Println(p3)
	//方法3：值列表初始化
	p4 := person{
		"机关炮",
		22,
	}
	fmt.Println(p4)

}

//问题3、为什么要有构造函数？
type person struct {
	name string
	age  int
}

func newPerson(x string, y int) person { //返回值类型
	//别人调用我，我能给他一个person类型的变量
	return person{
		name: x,
		age:  y,
	}
}

/*
func newPerson(x string, y int) *person { // 返回指针类型
	var q = &person{
		name: x,
		age:  y,
	}
	return q
}
*/
