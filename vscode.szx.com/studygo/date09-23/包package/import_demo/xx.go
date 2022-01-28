package main

import (
	"fmt"

	shizhenxing "vscode.szx.com/studygo/date09-23/包package/calc"
)

var x = 100

const pi = 3.1415

func init() {
	fmt.Println("自动执行!")
	fmt.Println(x, pi)
}

func main() {
	ret := shizhenxing.Add(10, 20)
	fmt.Println(ret)
}
