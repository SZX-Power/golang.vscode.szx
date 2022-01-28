package main

import "fmt"

//同一个结构体可以实现多个接口
//接口还可以嵌套
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}
type eater interface {
	eat(string)
}
type cat struct {
	name string
	feet int
}

//cat实现了mover接口
func (c cat) move() {
	fmt.Println("走猫步。。。")
}

//cat实现了eater接口
func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
	var a mover
	var b eater

	c := cat{"Tom", 4}
	d := cat{"Jack", 2}
	a = c
	b = d
	fmt.Println(a)
	fmt.Println(b)
}

/*
	接口和类型的关系:
	多个类型可以实现同一个接口
	一个类型可以实现多个接口
*/
