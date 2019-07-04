package main

import "fmt"

//iota从0开始，会自增
const (
	read byte = 1 << iota //0001
	write //此时iota=1,左移一位  0010
	exec
	freeze
)


func main(){

	a := read | write | freeze
	b := read | freeze | exec
	c := a &^ b //相当于 a ^ read ^ freeze = write 的值
	fmt.Printf("%04b &^ %04b = %04b\n", a, b, c)
}
