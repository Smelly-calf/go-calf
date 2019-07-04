package main

import "fmt"

// 位移右操作数必须是无符号整形 ，左边需要显式声明类型
func main()  {
	fmt.Printf("%T, %v \n", 1.0, 1.0)
	a := 1.0 << 3
	fmt.Printf("%T, %v \n", a, a)

	//var s uint = 8
	//b := 1.0 << s
	//fmt.Printf("%T, %v \n", b, b)


}
