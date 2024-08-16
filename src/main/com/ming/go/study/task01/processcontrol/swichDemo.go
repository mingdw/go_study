package main

import "fmt"

func main() {
	fmt.Println(getWenkendStr(1))
	fmt.Println(getWenkendStr(5))
}

func getWenkendStr(a int) string {
	switch a {
	case 1:
		return "星期一"
	case 2:
		return "星期二"
	case 3:
		return "星期三"
	case 4:
		return "星期四"
	case 5:
		return "星期五"
	case 6:
		return "星期六"
	case 7:
		return "星期七"
	default:
		return ""
	}
	return ""
}
