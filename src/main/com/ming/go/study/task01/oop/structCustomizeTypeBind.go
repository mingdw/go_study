package oop

import "fmt"

/*
*
自定义属性绑定函数
*/
func StructCustomizeTypeBindMain() {
	a := 30
	var is integer = 20
	is.setInteger(a)
	fmt.Println(is)
}

type integer int

func (i *integer) setInteger(a int) {
	*i = integer(a)
}
