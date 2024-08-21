package routine

import (
	"fmt"
	"strconv"
	"time"
)

/*
*
协程简单调用，如果主线程结束，携程也会结束
*/
func RoutineMain() {
	//每隔一秒程序交替打印内容，打印10次
	fmt.Println("********************")
	go print01()
	print02()
}

func print01() {
	for i := 1; i <= 10; i++ {
		fmt.Println("print01 " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func print02() {
	for i := 1; i <= 10; i++ {
		fmt.Println("print02 " + strconv.Itoa(i))
		/*if i==2 {
			return
		}*/
		time.Sleep(time.Second)
	}
}
