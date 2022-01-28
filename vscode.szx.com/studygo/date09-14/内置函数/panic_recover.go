package main

import "fmt"

/*
	1、recover()必须搭配defer使用。
    2、defer一定要在可能引发panic的语句之前定义。
*/

func funcA() {
	fmt.Println("a")
}

func funcB() {
	// 刚刚打开数据库连接
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接。。。")
	}()
	panic("出现了严重的错误！！！") // 程序崩溃退出
	fmt.Println("b")     //该语句永远执行不到 unreachable code
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
