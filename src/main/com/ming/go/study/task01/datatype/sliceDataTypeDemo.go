package main

import "fmt"

/*
*
切片相关操作
*/
func main() {
	//切片相比较数组，切片数据类型能够动态扩容，基于动态数组实现
	// 切片的定义
	var sl []string
	fmt.Println("初始化切片： ", sl)
	fmt.Println("初始化切片大小： ", len(sl))

	var sl2 []string = []string{"a", "b", "c"}
	fmt.Println("初始化切片2： ", sl2)
	sl2 = append(sl2, "d")
	fmt.Println("切片增加元素： ", sl2)

	//make方式定义切片
	ints := make([]int, 2, 3)
	fmt.Println("ints 是否为空：", ints == nil)
	fmt.Println("ints的值：", ints)
	ints = append(ints, 11)
	fmt.Println("ints的值：", ints)

	var ints2 []string
	fmt.Println("ints2 是否为空：", ints2 == nil)
	fmt.Println("ints2：", ints2)

	names := [4]string{ //定义了一个字符串数组
		"尹正杰",
		"百度",
		"谷歌",
		"FQ",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	//切片的删除本质上是切片的截取
	//删除最后一个元素
	newslice := names[:len(names)-1]
	fmt.Println("删除最后一个元素以后新值： ", newslice)

}
