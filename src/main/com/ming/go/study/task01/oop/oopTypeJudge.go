package oop

import "fmt"

// golang的类型判断：在golang中，.()被称为类型断言，可以将一个interface{}类型的变量转换成其它类型，例如字符串类型、整数类型等等
func OOPTypeJudge() {
	var w interface{} = Woman{name: "zahngsan"}
	t, ok := w.(Person)
	fmt.Println(t, ok)

	var g interface{} = God{age: 1}
	t2, ok2 := g.(Person)
	fmt.Println(t2, ok2)
	fmt.Println("*********************")
}

type Person interface {
	GetName() string
}

type Woman struct {
	name string
}

func (w Woman) GetName() string {
	return w.name
}

type God struct {
	age int
}

func (g God) GetAge() int {
	return g.age
}
