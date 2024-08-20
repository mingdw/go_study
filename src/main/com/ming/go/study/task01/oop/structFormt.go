package oop

import "fmt"

// 面向对象结构体的转换，必须属性相同
type Students struct {
	Age int
}

type Class struct {
	Age int
}

func StructFormatMain() {
	var stu = Students{1}
	var cla = Class{2}
	stu = Students(cla)
	fmt.Println(stu)
	fmt.Println(cla)
}
