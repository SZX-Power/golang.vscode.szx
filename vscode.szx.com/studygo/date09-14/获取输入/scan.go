package main

import "fmt"

//Go语言fmt包下有fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，可以在程序运行过程中从标准输入获取用户的输入。

//获取用户输入
func main() {
	// var s string
	// fmt.Scan(&s)
	// fmt.Println("用户输入的内容是：", s)

	var (
		name  string
		age   int
		class string
	)
	// fmt.Scanf("%s %d %s\n", &name, &age, &class)
	// fmt.Println(name, age, class)

	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)

}
