package main

import (
	"fmt"
	"os"
)

//学生管理系统

var smr studentMgr //声明一个全局的变量学生管理对象smr
//菜单函数
func showMenu() {
	fmt.Println("-------------Welcome SMS!------------")
	fmt.Println(`
	1、查看所有学生
	2、添加学生
	3、修改学生
	4、删除学生
	5、退出	
	`)
}

func main() {
	smr = studentMgr{ // 修改的全局的那个变量
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		//等待用户输入选项
		fmt.Print("请输入序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是：", choice)

		switch choice {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudents()
		case 3:
			smr.editStudents()
		case 4:
			smr.deleteStudents()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("无效选项，请输入正确的序号")
		}
	}

}
