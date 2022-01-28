package main

import (
	"fmt"
)

//为什么要有类型断言呢？ 我想知道空接口接收的值具体是什么？
// 类型断言

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了！")
	} else {
		fmt.Println("传进来的是一个字符串：", str)
	}
}

// 类型断言2
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串：", t)
	case int:
		fmt.Println("是一个int：", t)
	case int64:
		fmt.Println("是一个int64：", t)
	case bool:
		fmt.Println("是一个布尔类型：", t)
	}
}

func main() {
	assign(100)
	assign2(true)
	assign2(int64(200))
	assign2("哈哈哈")
}

/*
	关于接口需要注意的是，只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。
	不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。
*/
