package processcontrol

import (
	"fmt"
)

func ForMain() {
	fmt.Println("111")
	for1()
	for2()
	for3()
}

func for1() {
	i := 0
	slice := []int{}
	//直接循环
	for i < 10 {
		slice = append(slice, i) // 添加另一个切片的所有元素
		i++
	}
	fmt.Println(slice)
}

func for2() {
	slice := []int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 4}
	//range循环
	for i := range slice {
		fmt.Println(i)
	}

}

func for3() {
	slice := []int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 4}
	//range循环
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

}
