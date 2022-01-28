package main

import "fmt"

// 构造函数
type person struct {
	name, gender string
	age          int
	id           int
	y            bool
}

type dog struct {
	name string
	age  int
}

func newDog(name string, age int) dog {
	return dog{
		name: name,
		age:  age,
	}
}

/*
	构造函数：约定成俗用new开头
	返回的是结构体还是结构体指针
	当结构体比较大的时候尽量使用结构体指针，减少程序的开销
*/
func newPerson(name, gender string, age, id int) *person {
	return &person{
		name:   name,
		gender: gender,
		age:    age,
	}
}

func main() {
	p1 := newPerson("魏无羡", "男", 22, 1)
	p2 := newPerson("小樱", "女", 18, 2)
	fmt.Println(p1, p2)

	p3 := newDog("旺财", 2)
	p4 := newDog("小黑", 3)
	fmt.Println(p3, p4)
}
