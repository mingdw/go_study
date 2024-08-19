package datatype

import "fmt"

/*
*
数组数据类型相关操作
*/
func ArrayMain() {
	//数组大小一确定没法更改大小
	var str [3]string = [3]string{"a", "b", "c"}
	str2 := [3]string{"d", "e", "f"}
	fmt.Println("打印当前数组： ", str)
	fmt.Println("打印当前数组： ", str2)
	//获取当前数组的大小
	fmt.Println("打印当前数组的大小： ", len(str2))
	str[0] = "2"
	fmt.Println("改变数组的值： ", str)
	fmt.Println("获取某一个数组下表的值： ", str[1])
	for i := range str {
		fmt.Println("循环数组： ", i)
	}
}
