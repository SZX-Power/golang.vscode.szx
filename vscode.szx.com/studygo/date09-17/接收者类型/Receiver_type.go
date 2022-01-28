package main

import "fmt"

//接收者类型：值类型的接收者   指针类型的接收者

/*什么时候应该使用指针类型接收者
需要修改接收者中的值
接收者是拷贝代价比较大的大对象
保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/

type person struct {
	name string
	age  int
}

// 使用值接收者：（传拷贝进去）
func (p person) guonian() {
	p.age++
}

//使用指针接收者：（传内存地址进去）
func (p *person) zhenguonian() {
	p.age++
}

func (p *person) dream() {
	fmt.Println("不上班也能挣钱！")
}
func main() {
	var per person
	per.name = "元素"
	per.age = 18
	fmt.Println(per)

	per.guonian()
	fmt.Println(per.age)

	per.zhenguonian()
	fmt.Println(per.age)

	per.dream()

}
