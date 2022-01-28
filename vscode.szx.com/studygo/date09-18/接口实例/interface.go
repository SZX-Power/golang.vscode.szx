package main

import "fmt"

//接口示例    不管什么牌子的车，都能跑

//定义一个car接口类型
//不管是什么结构体  只要有run方法都是car类型
type car interface {
	run()
}

type Ferrari struct {
	brand string
}

func (F Ferrari) run() {
	fmt.Printf("%s速度70迈~\n", F.brand)
}

type Porsche struct {
	brand string
}

func (P Porsche) run() {
	fmt.Printf("%s速度是80迈~\n", P.brand)
}

//drive函数接收一个car类型的变量
func drive(c car) {
	c.run()
}

func main() {
	var F1 = Ferrari{
		brand: "法拉利",
	}
	var P1 = Porsche{
		brand: "保时捷",
	}

	drive(F1)
	drive(P1)

}
