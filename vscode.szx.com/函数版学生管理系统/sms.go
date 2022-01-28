package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够查看\新增\删除学生
*/

type student struct {
	id   int64
	name string
}

var allstudent map[int64]*student //变量声明

// newStudent 是student类型的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func studentAll() {
	//把所有学生打印出来
	for k, v := range allstudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
}

func studentAdd() {
	//向allstudent中添加一个新的学生
	//1、创建一个新学生
	//2、追加到allstudent这个map中
	//1.1获取用户输入
	var id int64
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)
	var name string
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	// if id == allstudent[id] || name == allstudent {
	// 	fmt.Println("提示：此学生已存在,请勿重复添加！")
	// } else {
	//1.2造学生（调用构造函数）
	newStu := newStudent(id, name)
	//2.追加到allstudent这个map中
	allstudent[id] = newStu
	//}
}

func studentDel() {
	// 1、请用户输入要删除的学生的序号
	// 2、去allstudent这个map中根据学号删除对应的键值对
	var (
		Delid int64
	)
	fmt.Print("请输入要删除学生的学号：")
	fmt.Scanln(&Delid)

	delete(allstudent, Delid)
}

func main() {
	allstudent = make(map[int64]*student, 48) //初始化（开辟内存空间）
	for {
		// 1、打印菜单
		fmt.Println("---欢迎来到学生管理系统---")
		menu :=
			`
		1、查看所有学生信息
		2、增加学生信息
		3、删除学生信息
		4、退出
		`
		fmt.Println(menu)
		fmt.Print("请输入你的选择：")
		// 2、等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项！", choice)
		// 3、执行对应的函数
		switch choice {
		case 1:
			studentAll()
		case 2:
			studentAdd()
		case 3:
			studentDel()
		case 4:
			os.Exit(1) //退出
		default:
			fmt.Println("无效选项，请输入正确的序号")
		}
	}

}
