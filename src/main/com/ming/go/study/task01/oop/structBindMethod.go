package oop

import "fmt"

type EmallOrder struct {
	orderDate   string // 订单时间
	orderStatus int    //订单状态
	orderAmount int    //订单金额
	payType     string
}

func (emallOrder EmallOrder) getOrderDate() string {
	return emallOrder.orderDate
}

func (emallOrder EmallOrder) setOrderDate(orderDates string) {
	emallOrder.orderDate = orderDates
	fmt.Println("this order: ", emallOrder)
}

func StructBindMain() {
	var order = EmallOrder{}
	order.setOrderDate("20220981")
	fmt.Println("order: ", order)
}
