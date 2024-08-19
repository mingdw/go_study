package method

import "fmt"

/*
闭包
*/
func ClosePackageMain() {
	f := getSum()
	fmt.Println(f(1, 2))
	fmt.Println(f(3, 4))
	fmt.Println(f(5, 6))

}

func getSum() func(a int, b int) int {
	var s int = 0
	return func(a int, b int) int {
		s = s + a + b
		return s
	}
}
