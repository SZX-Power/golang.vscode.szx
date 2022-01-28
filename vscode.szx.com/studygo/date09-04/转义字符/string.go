package main

import (
	"fmt"
	"strings"
)

func main() {
	// \本来是具有特殊含义的，我应该告诉程序我写的\就是一个单纯的\
	path := "E:\\Go\\goDemoLearn\\src\\code.szx.com\\studygo"
	fmt.Println(path)
	const path01 = "\"E:\\Go\\goDemoLearn\\src\\code.szx.com\\studygo\""
	var path02 = "'E:\\Go\\goDemoLearn\\src\\code.szx.com\\studygo'"
	fmt.Println(path01, path02)

	s := "I'm ok"
	fmt.Println(s)

	//多行的字符串
	s2 := `
	世情薄，
	人情恶，
	雨送黄昏花易落。
	`
	fmt.Println(s2)
	s3 := `E:\Go\goDemoLearn\src\code.szx.com\studygo`
	fmt.Println(s3)

	//字符串相关操作
	fmt.Println(len(s3))

	//字符串拼接
	name := "理想"
	world := "power!"
	//方法一
	ss := name + world
	fmt.Println(ss)
	//方法二
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)
	//方法三
	fmt.Printf("%s%s", name, world)
	fmt.Println("----------------------")

	//分割
	aa := strings.Split(s3, "\\")
	fmt.Println(aa)
	//判断是否包含
	fmt.Println(strings.Contains(ss, "理性"))
	fmt.Println(strings.Contains(ss, "理想"))
	//判断前缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	//判断后缀
	fmt.Println(strings.HasSuffix(ss, "理想"))

	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))

	//拼接
	fmt.Println(strings.Join(aa, "+"))

}
