package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

/*
整体运行逻辑顺序：
	1. x:=1
	2. y:=2
	3. defer calc("AA",1,calc("A",1,2))
	4. calc("A",1,2)  //A,1,2,3
	5. defer calc("AA",1,3)
	6. x =10
	7. defer calc("BB",10,calc("B",10,2))
	8. calc("B",10,2)  //"B",10,2,12
	9. defer calc("BB",10,12)
	10. y=20
	11. defer calc("BB",10,12) //"B",10,12,22
	12. defer calc("AA",1,3)  // "AA"1,3,4
*/
