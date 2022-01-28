package main

import "fmt"

//接口的实现

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步....")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)

}

func (c chicken) move() {
	fmt.Println("鸡动！")
}
func (c chicken) eat(food string) {
	fmt.Printf("鸡吃%s！\n", food)
}

func main() {
	var a animal // 定义了一个接口类型的变量
	fmt.Printf("%T\n", a)
	bc := cat{
		name: "淘气",
		feet: 4,
	}
	a = bc
	a.eat("小黄鱼")
	fmt.Printf("%T\n", a)
	fmt.Println(a)

	kfc := chicken{
		feet: 2,
	}

	a = kfc
	fmt.Printf("%T\n", a)

}
