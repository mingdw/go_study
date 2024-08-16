package main

import "fmt"

func main() {
	var a = true
	if a {
		fmt.Println("输出： true")
	}

	caishu(15)
}

func caishu(a int) bool {
	if a < 10 {
		return false
	} else {
		return true
	}
}
