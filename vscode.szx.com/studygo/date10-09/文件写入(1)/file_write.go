package main

import (
	"fmt"
	"os"
)

//打开文件内容
//0100  0000
func main() {
	fileObj, err := os.OpenFile("./szx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	//write  写入字节切片数据
	fileObj.Write([]byte("zhoulin menbi le !\n"))

	//writeString  直接写入字符串数据
	fileObj.WriteString("周林懵逼了！")
	fileObj.Close()
}
