package method

import "fmt"

func DeferMain() {
	fmt.Println("result： ", sum(10, 20))
}

func sum(a int, b int) int {
	fmt.Println("sum start")
	//defer 关键字的代码部分不会立即执行，会等本次方法执行完以后再执行
	defer fmt.Println("num1 ", a)
	defer fmt.Println("num2 ", b)
	var s = a + b
	fmt.Println("sum result: ", s)
	return s
}
