package main

import "fmt"

//for循环
func main() {
	//基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("-----------------------")
	//变种1
	var y = 5
	for ; y < 10; y++ {
		fmt.Println(y)
	}
	fmt.Println("-----------------------")
	//变种2
	var x = 5
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("-----------------------")
	//无线循环
	// for {
	// 	fmt.Println("123")
	// }

	//for range 键值循环
	s := "Hello沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
