package oop

import (
	"../datatype"
	"encoding/json"
	"fmt"
)

type EmallOrder struct {
	orderDate   string // 订单时间
	orderStatus int    //订单状态
	orderAmount int    //订单金额
	payType     string
}

func (e EmallOrder) getOrderDate() string {
	return e.orderDate
}

func (e *EmallOrder) setOrderDate(orderDates string) {
	e.orderDate = orderDates
}

func (e EmallOrder) getOrderorderStatus() int {
	return e.orderStatus
}

func (e *EmallOrder) setOrderorderStatus(orderStatus int) {
	e.orderStatus = orderStatus
}

func (e EmallOrder) getOrderAmount() int {
	return e.orderAmount
}

func (e *EmallOrder) setOrderAmount(amount int) {
	e.orderAmount = amount
}

// 如果有自定义的String方法
func (s *EmallOrder) String() string {
	jsonStr, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(jsonStr)
	return string(jsonStr)
}

func StructBindMain() {
	order := EmallOrder{
		orderDate:   "2024-08-19",
		orderAmount: 12334,
		orderStatus: 10,
		payType:     "Wechat",
	}
	order.setOrderDate("2024-08-19")
	order.setOrderAmount(100)
	order.setOrderorderStatus(0)
	fmt.Println("order: ", order)

	//跨包导入结构体,同一层直接写实体名称
	stu := Student{
		Name: "zhansa",
		Age:  13,
	}
	fmt.Println("跨包导入： ", stu)

	compnay := datatype.Company{
		Name:  "sss",
		Age:   123,
		Phone: "eqwee",
	}
	fmt.Println("跨包导入2： ", compnay)
}
