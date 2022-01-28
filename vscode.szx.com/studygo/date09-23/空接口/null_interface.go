package main

import "fmt"

// 空接口，因为空接口可以存储任意类型值的特点，所以空接口在Go语言中的使用十分广泛。

// interface：关键字
// interface{}：空接口类型

//空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
func main() {
	var m map[string]interface{}
	m = make(map[string]interface{}, 16)
	m["name"] = "周林"
	m["age"] = 18
	m["merried"] = true
	m["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m)

	show(false)
	show(nil)
	show(m)
}
