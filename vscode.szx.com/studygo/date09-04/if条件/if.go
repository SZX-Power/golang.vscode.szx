package main

import "fmt"

func main() {
	age := 23                //两个age变量的作用域不同
	if age := 19; age > 18 { //此age只作用在if循环内
		fmt.Println("成年啦！")
	} else if age > 35 {
		fmt.Println("而立之年")
	} else {
		fmt.Println("成长中~")
	}
	fmt.Println(age)

	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

}
