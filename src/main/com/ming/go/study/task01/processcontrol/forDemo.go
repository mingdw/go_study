package main

import (
	"fmt"
	"strconv"
)

func main() {
	For1()
	for2()
	for3()
}

func For1() {
	ar := make([]string, 10, 15)
	//计数增加循环
	for i := 0; i < len(ar); i++ {
		ar[i] = "num:" + strconv.Itoa(i)
	}
	fmt.Println(ar)
}

func for2() {
	ar := make([]string, 10, 15)
	var i int
	//条件循环
	for i < 10 {
		ar[i] = "num:" + strconv.Itoa(i)
		i++
	}
	fmt.Println(ar)
}

func for3() {
	ar := make([]string, 10, 15)
	//range循环
	for i := range ar {
		ar[i] = "num:" + strconv.Itoa(i)
	}
	fmt.Println(ar)
}
