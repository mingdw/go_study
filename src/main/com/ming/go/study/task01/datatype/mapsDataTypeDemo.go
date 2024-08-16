package main

import "fmt"

/*
*
maps数据结构
*/
func main() {
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 7
	fmt.Println("map:", m)

	m2 := map[string]int{"age": 11, "sex": 1}
	fmt.Println("map2:", m2)

	//新增
	m2["phone"] = 1871203
	fmt.Println("add value map2:", m2)

	//删除key
	delete(m2, "phone")
	fmt.Println("delete key map2:", m2)

	//是否包含某个key
	value, ok := m2["age"]
	if ok {
		fmt.Println("map exsist key, value is :", value)
	} else {
		fmt.Println("map not exsist key!")
	}

	//遍历
	for s := range m2 {
		fmt.Println("loop map，key is: %c, value is: %c", s, m2[s])
	}

	//覆盖原值
	m2["sex"] = 10
	fmt.Println("overrload old value, the new value is: ", m2)

	//清空
	clear(m2)
	//date_type := DATE_TYPE

}
