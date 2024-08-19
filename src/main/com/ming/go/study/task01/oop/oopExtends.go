package oop

import "fmt"

type Animal struct {
	Anum   string
	Age    int
	Weight float32
}

func (an *Animal) Eating() {
	fmt.Println("执行动物吃饭的方法。。。。。。。。。。")
}

func (an *Animal) Display() {
	fmt.Println("动物展示。。。", *an)
}

type Cart struct {
	Animal
	Name string
}

func (cart *Cart) CatchMourse() {
	fmt.Println("猫能够捕捉老鼠", cart.Name)
}

type Dog struct {
	Animal
	Dname string
}

func (dog *Dog) CatchCart() {
	fmt.Println("狗拿猫", dog.Dname)
}

func OOPExtendsMain() {
	var cart = Cart{}
	cart.Age = 12
	cart.Anum = "C009122"
	cart.Name = "橘猫"
	fmt.Println("cart info: ", cart)
	cart.CatchMourse()
	cart.Display()

	var dog = Dog{}
	dog.Dname = "金毛"
	fmt.Println("dog info: ", dog)
	dog.CatchCart()
	dog.Display()
}
