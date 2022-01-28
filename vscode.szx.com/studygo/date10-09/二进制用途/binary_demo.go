package main

import "fmt"

//二进制实际用途
const (
	eat   int = 4
	sleep int = 2
	run   int = 1
)

/*111
左边的1表示吃饭		100
中间的1表示睡觉		010
右边的1表示跑步		001
*/

func f(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	f(eat | sleep)
	f(sleep | run)
	f(eat | sleep | run)
}
