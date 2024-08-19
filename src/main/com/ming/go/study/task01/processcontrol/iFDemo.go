package processcontrol

import "fmt"

func IfMain() {

	fmt.Println("15和16哪个大", getMax(15, 16), "大")
}

func getMax(a int, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return a
	} else {
		return b
	}
}
