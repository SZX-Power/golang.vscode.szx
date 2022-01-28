package main

import (
	"fmt"
	"io"
	"os"
)

//打开文件  (按字节读)

func main() {
	fileObj, err := os.Open("./file_open.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()

	//读文件
	//var tmp = make([]byte,128)  //指定读的长度
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了！")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n == 0 { //n<128
			return
		}
	}

}
