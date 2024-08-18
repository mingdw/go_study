package exception

import (
	"errors"
	"fmt"
)

/*
*
自定义异常
*/
func CusException() {
	r, err := add2(5, 0)
	if err != nil {
		//如果希望程序终止，调用panic
		panic(err)
	}
	fmt.Println("计算结果， ", r, "是否有异常： ", err)
}

func add2(a int, b int) (r int, err error) {
	fmt.Println("除数a： ", a)
	fmt.Println("被除数b： ", b)
	//假设b==0，抛出自定义异常
	if b == 0 {
		return -1, errors.New("被除数不能为0")
	}
	return a / b, nil
}
