package main

import (
	"./method"
	"fmt"
)

/*
*
go的函数处理
*/
func main() {
	method.DeferMain()
	method.SystemFuncMain()
	method.DateMain()
	fmt.Println("********************")
}
