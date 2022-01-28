package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//文件写入（bufio.NewWriter , ioutil.WriteFile）

func bufioWriter() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	defer fileObj.Close()
	// 创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	// 将数据先写入缓存
	wr.WriteString("hello沙河\n")
	// 将缓存中的内容写入文件
	wr.Flush()
}

func ioutilWriter() {
	str := "hello,沙河。冲啊！"
	err := ioutil.WriteFile("./yy.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("write file failed,err:%v", err)
		return
	}
}

func main() {
	bufioWriter()
	ioutilWriter()

}
