package main

import "fmt"

//交换a,b地址所指向的值
func swap(a, b *int) {
	//将a指向的变量值赋值给b指针指向的变量值，将b指针指向的变量值赋值给a指针指向的变量值
	*b, *a = *a, *b
}

func main() {
	a, b := 3, 4
	//传入a和b的地址
	swap(&a, &b)
	fmt.Println(a, b)

	var c int = 2
	//定义一个指针pc指向变量c的地址
	var pc *int = &c
	//将pc指针指向的变量赋值为3
	*pc = 3
	//打印c的值，发现c的值由2变为3了
	fmt.Println(c)
}
