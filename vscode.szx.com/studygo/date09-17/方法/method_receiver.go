package main

import "fmt"

//方法

//标识符：变量名 函数名 类型名  方法名
//Go语言中如果标识符首字母是大写的，就表示对外部包可见（暴露的、公有的）

type dog struct {
	name string
}

//构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪~", d.name)
}

func main() {
	d1 := newDog("周林")
	d1.wang()

}
