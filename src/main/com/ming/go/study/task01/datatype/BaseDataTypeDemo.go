package main

import (
	"fmt"
	"strings"
)

/*
*
基本数据类型
*/
func main() {
	// go 的基本数据类型

	//1、整数类型：有符号整数和无符号整数
	var a int = -12
	var b uint = 12
	var c int
	fmt.Println("有符号整数:", a)
	fmt.Println("无符号整数:", b)
	fmt.Println("int默认值:", c)

	//2.浮点
	var f1 float32 = 12.3
	var f2 float32
	fmt.Println("浮点类型:", f1)
	fmt.Println("浮点类型默认值:", f2)

	//3.布尔
	var bl bool = true
	var bl2 bool
	fmt.Println("布尔类型:", bl)
	fmt.Println("布尔类型默认值:", bl2)

	//4. 字符串
	var s1 string = "这是一个字符串asd 01"
	var s2 string
	fmt.Println("字符串类型:", s1)
	fmt.Println("字符串类型默认值:", s2)
	//一个中文3个字节，一个字节1个长度
	fmt.Println("字符串长度:", len(s1))

	//字符串常见操作
	//4.1 字符串拼接
	apend := s1 + s2
	fmt.Println("字符窜拼接结果:", apend)

	//4.1 字符串是否包含某个字符
	str := "Hello World go 中"
	found := strings.Contains(str, "World")
	found2 := strings.Contains(str, "中")
	fmt.Println("是否包含： ", found) // 输出: true
	fmt.Println("是否包含", found2)  // 输出: true

	//4.1 字符串是否包含某个字符
	str2 := " Hello World go 中"
	found3 := strings.Replace(str2, "World", "City", -1)
	fmt.Println(found3)
	fmt.Println(found2) // 输出: true
	fmt.Println("字符串分割：", strings.Split(str2, " "))
	fmt.Println("字符串转大写：", strings.ToUpper(str2))
	fmt.Println("字符串转小写：", strings.ToLower(str2))
	fmt.Println("字符串去掉首尾空格：", strings.TrimSpace(str2))
	var sn string = `
					func main() {
						str := "   Hello World   "
						result := strings.TrimSpace(str)
						fmt.Println(result) // 输出: Hello World
					}
`
	fmt.Println("字符串多行展示：", strings.TrimSpace(sn))
}
