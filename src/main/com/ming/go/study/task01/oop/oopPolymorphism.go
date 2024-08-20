package oop

import "fmt"

// golang面向对象的多态实现
func PolymorphismMain() {
	var s Shape
	r := Rectange{width: 3, height: 4}
	s = r
	fmt.Println("正方形的面积： ", s.Area())

	var s2 Shape
	c := Circle{radius: 3}
	s2 = c
	fmt.Println("圆形的面积： ", s2.Area())
	fmt.Println("***********动态计算面积***********")
	fmt.Println("正方形的面积： ", GetArea(r))
	fmt.Println("圆形的面积： ", GetArea(c))

}

// 定义形状接口
type Shape interface {
	//计算面积的方法
	Area() float64
}

// 正方形的结构体
type Rectange struct {
	//正方形属性 宽
	width float64
	//正方形属性 高
	height float64
}

// 圆 的结构体
type Circle struct {
	//圆的属性 半径
	radius float64
}

// 给正方形绑定计算面积的方法
func (r Rectange) Area() float64 {
	return r.width * r.height
}

// 给圆形绑定计算面积的方法
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// 动态计算面积
func GetArea(s Shape) float64 {
	switch x := s.(type) {
	case Circle:
		return x.Area()
	case Rectange:
		return x.Area()
	default:
		return 0
	}
}
