package main

import "fmt"

//switch  简化大量的判断（一个变量和具体的值做比较）
func main() {
	var n = 3
	if n == 1 {
		fmt.Println("大拇指")
	} else if n == 2 {
		fmt.Println("食指")
	} else if n == 3 {
		fmt.Println("中指")
	} else if n == 4 {
		fmt.Println("无名指")
	} else if n == 5 {
		fmt.Println("小拇指")
	} else {
		fmt.Println("无效的数字")
	}

	//switch简化上面的代码
	m := 1
	switch m {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的数字")
	}

	//go语言规定每个switch只能有一个default分支
	//一个分支可以有多个值，多个case值中间使用英文逗号分隔
	switch b := 7; b {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(b)

	}
}
