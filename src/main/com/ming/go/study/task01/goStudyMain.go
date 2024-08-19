package main

import (
	"./datatype"
	"./method"
	"./processcontrol"
	"fmt"
)

func init() {
	fmt.Println("init.............")
}

func main() {
	fmt.Println("***********************hello world golang ***********************")
	fmt.Println("hellow world go!")
	fmt.Println("*********************** datatype ***********************")
	datatype.ArrayMain()
	datatype.BaseMain()
	fmt.Println(datatype.DATE_TYPE)
	datatype.MapsMain()
	datatype.PointMain()
	datatype.SliceMain()
	fmt.Println("*********************** method ***********************")
	method.MethodMain()
	method.ClosePackageMain()
	fmt.Println("*********************** processcontrol ***************")
	processcontrol.ForMain()
	processcontrol.IfMain()
	processcontrol.SwichMain()
}
