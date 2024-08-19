package exception

import "fmt"

/*
*
程序异常捕获，defer + recover（）
*/
func ExceptionMain() {
	fmt.Println("除数计算如下a b")
	r := add(5, 0)
	fmt.Println("除数计算结果： ", r)

}

func add(a int, b int) int {
	fmt.Println("除数a： ", a)
	fmt.Println("被除数b： ", b)

	return a / b
}
