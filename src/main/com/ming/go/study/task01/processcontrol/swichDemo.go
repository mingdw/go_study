package processcontrol

import "fmt"

func SwichMain() {
	fmt.Println("16是星期： ", getWenkendStr(17))
}

// 给定数字，计算是星期几
func getWenkendStr(a int) string {
	var b = a % 7
	switch b {
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
		return "星期天"
	default:
		return ""
	}
	return ""
}
