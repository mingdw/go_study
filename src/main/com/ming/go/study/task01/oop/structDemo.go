package oop

import "fmt"

type Student struct {
	Name      string //姓名
	Age       int    //年龄
	ClassRoom string //班级
	Wechat    string //微信
	Phone     string //电话
}

func StruMain() {
	initStudent1()
	initStudent2()
	initStudent3()
}

func initStudent1() {
	var stu01 = Student{"张三", 11, "三年二班", "wecht001", "18710192763"}
	fmt.Println("init01: ", stu01)
}

func initStudent2() {
	var stu02 *Student = new(Student)
	stu02.Name = "赵四"
	stu02.Age = 15
	stu02.ClassRoom = "四年5班"
	fmt.Println("init02: ", *stu02)
}

func initStudent3() {
	var stu03 *Student = &Student{}
	stu03.Name = "王五"
	stu03.Age = 17
	stu03.ClassRoom = "四年5班"
	fmt.Println("init03: ", *stu03)
}
