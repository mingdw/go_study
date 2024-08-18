package method

import (
	"fmt"
)

/*
函数定义
*/
func MethodMain() {
	x, y := del(12, 15)
	fmt.Print(x, y)
	fmt.Println(del2(12, 13))
	fmt.Println(del4(1, 2, 34))
	fmt.Println(del6(12, 34))
	fmt.Println(del7(1, 2))
}

func del0(a int, b int) int {
	return a * b
}

/*
*
返回多个参数
*/
func del(a int, b int) (x int, y int) {
	x = a + b
	y = a - b
	return x, y
}

/*
如果返回一个参数，返回列表省略括号
*/
func del2(a int, b int) int {
	x := a + b
	return x
}

func del4(args ...int) int {
	var count = 0
	for i := range args {
		count = count + i
	}
	return count
}

// 函数作为参数调用
func del5(a int, b int, del0 func(a int, b int)) int {
	return a + b
}

/*
返回参数直接赋值
*/
func del6(a int, b int) (x int, y int) {
	x = a * b
	y = a + b
	return
}

// 如果希望函数只被使用一次，使用匿名函数
func del7(a int, b int) (x int, y int) {
	//匿名函数就是内部函数，定义的时候就被执行
	s := func(q int, w int) int {
		return a + b
	}(a, b)
	x = s + a + b
	y = a * b
	return
}
