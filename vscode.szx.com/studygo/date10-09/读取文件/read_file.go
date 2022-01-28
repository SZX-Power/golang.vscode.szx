package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//读取文件另外两种方法（Bufio,ioutil）

//bufio是在file的基础上封装了一层API，支持更多的功能。
//按行读取
func readFromFilebyBufio() {
	fileObj, err := os.Open("E:/Go/goWorkspace/src/vscode.szx.com/studygo/date10-09/打开文件/file_open.go")
	if err != nil {
		fmt.Printf("文件打开失败，err:%v", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 创建一个用来从文件中都内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		reader.ReadString('\n') //字符 按行读取
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读完了！")
			return
		}
		if err != nil {
			fmt.Printf("read line failed,err:%v", err)
			return
		}
		fmt.Print(line)
	}

}

// io/ioutil包的ReadFile方法能够读取完整的文件,只需要将文件名作为参数传入
func readFromFilebyIoutil() {
	ret, err := ioutil.ReadFile("E:/Go/goWorkspace/src/vscode.szx.com/studygo/date10-09/打开文件/file_open.go")
	if err != nil {
		fmt.Printf("read file失败，err:%v", err)
		return
	}
	fmt.Print(string(ret))
	//fmt.Print(ret)
}

func main() {
	readFromFilebyBufio()
	readFromFilebyIoutil()
}
