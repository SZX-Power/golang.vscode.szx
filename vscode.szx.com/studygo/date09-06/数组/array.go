package main

import "fmt"

// 数组  存放元素的容器
// 必须指定存放的元素类型和容量（长度）
// 数组的长度是数组类型的一部分， 所以例如如下a1和a2无法比较

type szx interface {
}

func main() {
	var a1 [3]bool // [true false true]
	var a2 [4]bool // [true true false false]
	fmt.Printf("a1:%T a2:%T\n0", a1, a2)

	// 数组初始化
	// 如果不初始化，默认元素都是零值（布尔值：false，整型和浮点型都是:0，字符串：""）
	fmt.Println(a1, a2)

	//go是强类型语言，但任何类型都可以转为interfalce{}类型。 所以用 []interface{}
	h := [...]szx{1, 2, 3.14, "hello", true}
	fmt.Println(h)

	//1、初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//2、初始化方式2
	b := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(b)
	//3、初始化方式3：根据索引来初始化
	a3 := [5]int{1, 2} //后面没有的，默认为0
	fmt.Println(a3)
	a4 := [5]int{0: 1, 4: 2} //根据索引序号
	fmt.Println(a4)

	//数组的遍历
	citys := [...]string{"青岛", "北京", "上海", "深圳"} //索引：0~3 citys[0],citys[1],citys[2],citys[3]
	// 1、根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//2、for range 遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	//多维数组
	//[[1 2] [3 4] [5 6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a11)

	//多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	// 数组是值类型
	b1 := [3]int{1, 2, 3} // [1 2 3]
	/*b1[0] = 100 这样也可以
	fmt.Println(b1)
	*/
	b2 := b1            // [1 2 3] Ctrl+C Ctrl+V =>把word文档从文件夹A拷贝到文件夹B
	b2[0] = 100         // b2:[100 2 3]
	fmt.Println(b1, b2) // b1:?
}
