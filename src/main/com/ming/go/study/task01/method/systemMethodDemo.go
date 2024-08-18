package method

import (
	"fmt"
	"strconv"
	"strings"
)

/*
*
系统函数相关
*/
func SystemFuncMain() {
	aboutStr()
}

/*
*
字符串相关
*/
func aboutStr() {
	s1 := "hello go中文" // 中文占3个字节
	fmt.Println(s1, "字节长度是： ", len(s1))

	//字符串的遍历
	for i, value := range s1 {
		fmt.Println("索引为:", i, "值是:", value)
	}

	//字符串转整数
	n, _ := strconv.Atoi("1232")
	fmt.Println("字符串转数字 用strconv.Atoi", n)

	//数字转字符串
	s2 := strconv.Itoa(1232)
	fmt.Println("字符串转数字 用strconv.Itoa", s2)

	//统计一个字符串中出现了几个指定的字符
	fmt.Println("golang is very good 中出现的o的次数是", strings.Count("golang is very good", "o"))

	//不区分大小写的比较
	fmt.Println("Go 和 go 字符串的比较： ", strings.EqualFold("go", "Go"))
	//区分大小写的比较
	fmt.Println("Go 和 go 字符串的比较： ", "Go" == "go")

	//首次出现在某个字符中首尾索引值,没有出现返回-1
	fmt.Println("首次出现在某个字符中首尾索引值", strings.Index("abcdedascd", "d"))

	//字符串的替换n=-1,全部替换，其它为替换的次数
	fmt.Println("golang 中的go 替换成 java：", strings.Replace("golang", "go", "java", -1))
	fmt.Println("golanggo 中的go 替换1次成 java：", strings.Replace("golanggo", "go", "java", 1))

	//字符串的分割
	fmt.Println("go-java-js-mysql按照-分割：", strings.Split("go-java-js-mysql", "-"))

	//大小写转换
	fmt.Println("go 转换  GO：", strings.ToLower("go"))
	fmt.Println("Go 转换  go：", strings.ToUpper("GO"))

	//字符串去掉空格
	fmt.Println(" 去掉空格：", strings.TrimSpace(" 去掉空格："))

	//去掉指定左右两边字符
	fmt.Println("——去掉指定左右两边字符——", strings.Trim("——去掉指定左右两边字符——", "——"))
}
