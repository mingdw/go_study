package datatype

import "fmt"

/*
*
指针
*/
func PointMain() {
	s1 := "这是中文字符串"
	i1 := 12
	s2 := "这是中文字符串"
	fmt.Println("字符串地址1：", &s1)
	fmt.Println("int地址：", &i1)
	fmt.Println("s1==s2?：", s1 == s2)

	var s1Point *string = &s1
	fmt.Println("指针地址：", s1Point)
	fmt.Println("指针实际代表的值：", &s1Point)
	fmt.Println("指针实际代表的值：", *s1Point)

	//通过指针改变所有的值
	var p1 int = 10
	var p2 int = 10
	var p3 int = 10
	var p4 int = 10

	fmt.Println("p1、p2、p3、p4指针分别为：", &p1, &p2, &p3, &p3)
	var pp1 *int = &p1
	var pp2 *int = &p2
	var pp3 *int = &p3
	var pp4 *int = &p4

	*pp1 = 20
	*pp2 = 20
	*pp3 = 20
	*pp4 = 20
	fmt.Println("p1、p2、p3、p4指修改指针地址后的值：", p1, p2, p3, p4)

}
