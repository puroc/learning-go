package main

import (
	"fmt"
	"math"
)

//在函数外部定义的变量是包内变量，不能使用:=定义变量
var (
	aa = 3
	bb = 4
)

//常量定义
const (
	a, b = 3, 4
)

//go语言自动为变量赋予初始值
func variableZeroValue() {
	var a int
	var b string
	fmt.Println(a, b)
}

//手动对变量赋予初始值
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "haha"
	fmt.Println(a, b, s)
}

//根据变量的初始值，推断变量的类型，声明变量时不需要加类型
func variableTypeDeduction() {
	var a, b, c, d = 8, 9, "xixi", true
	fmt.Println(a, b, c, d)
}

//定义变量时，省略var关键字
func variableShorter() {
	//在定义变量时，可以省略var关键字，改为:=对变量进行定义。达到简写的目的
	a, b, c, d := 8, 9, "xixi", true
	//但后续再对变量进行赋值时，只能写=号，不能写:=
	b = 10
	fmt.Println(a, b, c, d)
}

//强制类型转换
func typeConvert() {
	var a, b int = 3, 4
	var c int
	//在go语言中不会进行类型的自动转换，需要使用强制类型转换。
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//常量定义
func consts() {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//通过常量const实现枚举类型的定义，iota代表一个自增数列，从0开始
func enums() {
	//a=iota说明a=0，b、c、d、e依次递增
	const (
		a = iota
		b
		c
		d
		e
	)
	fmt.Println(a, b, c, d, e)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb)
	typeConvert()
	consts()
	enums()
}
