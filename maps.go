package main

import "fmt"

func main() {

	//定义一个map，并赋予初值。map的key不能是切片、map和func类型的
	m := map[string]string{
		"a": "1",
		"b": "2",
	}

	//定义一个空的map
	m2 := make(map[string]int)

	//定一个map，值为nil
	var m3 map[string]int

	fmt.Println(m, m2, m3)

	//遍历map，map的遍历是不保证顺序的
	for k, v := range m {
		fmt.Println(k, v)
	}

	//根据key获取map中的值
	fmt.Println(m["a"])

	//根据key获取map中的值时，返回的第二个值是false代表该值不存在，返回true代表该值存在
	v, exist := m["c"]
	if !exist {
		fmt.Println("not exist")
	} else {
		fmt.Println(v)
	}

	//删除元素
	delete(m, "b")
	v2, exist2 := m["b"]
	if exist2 {
		fmt.Println(v2)
	} else {
		fmt.Println("not exist")
	}

}
