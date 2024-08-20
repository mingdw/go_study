package method

import (
	"fmt"
	time "time"
)

/*
*
日期相关的函数
*/
func DateMain() {
	//获取当前时间
	nowtime := time.Now()
	fmt.Println("当前时间：", nowtime)
	fmt.Println("当前时间年：", nowtime.Year())
	fmt.Println("当前时间月：", nowtime.Month())
	fmt.Println("当前时间月（转换）：", int(nowtime.Month()))
	fmt.Println("当前时间日：", nowtime.Day())
	fmt.Println("当前时间时：", nowtime.Hour())
	fmt.Println("当前时间分：", nowtime.Minute())
	fmt.Println("当前时间秒：", nowtime.Second())

	//日期的格式化
	fmt.Println("当前时间按照yyyy-MM-dd HH:mm:ss 表示：", nowtime.Format("2006-01-02 15:04:05"))
}
