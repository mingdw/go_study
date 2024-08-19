package oop

import "fmt"

/*
golang面向对象的封装，
所谓封装就是将抽象出的字段和对应的操作封装在一起，数据被保护在内部，程序的其它包只能通过被授权的方法才能访问相关属性
*/

type person struct {
	Name   string
	age    int
	idCard string
}

func NewPereson(name string, age int, id string) *person {
	return &person{
		Name:   name,
		age:    age,
		idCard: id,
	}
}

func CreatPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age <= 150 {
		p.age = age
	} else {
		fmt.Println("年龄输入不合法")
	}

}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetIdCard(id string) {
	p.idCard = id

}

func (p *person) GetIdCard() string {
	return p.idCard
}

func OOPPackageMain() {
	p := CreatPerson("zhangsan ")
	p.SetAge(123)
	p.SetIdCard("4290016")
	fmt.Println("person: ", *p)
}
