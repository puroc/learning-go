package main

import (
	"fmt"
	"io/ioutil"
)

func if1() {
	const filename = "abc.txt"
	contents, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func if2() {
	const filename = "abc.txt"
	//if语句后面可以跟多个语句，语句之间用分号分隔。在if语句中定义的变量，这些变量的作用域就在这个if语句中。
	if contents, error := ioutil.ReadFile(filename); error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

//switch后面有表达式
//switch case会自动break，不用写break关键字。panic代表出错，会中断程序。
func case1(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupport operator:" + op)
	}
	return result
}

//switch后面没有表达式，在case后面写表达式
func case2(score int) string {
	grade := "NA"
	switch {
	case score < 60:
		grade = "D"
	case score < 70:
		grade = "C"
	case score < 80:
		grade = "B"
	case score < 90:
		grade = "A"
	default:
		grade = "E"
	}
	return grade
}

func main() {
	if1()
	if2()
	result := case1(1, 2, "+")
	fmt.Println(result)
	result2 := case2(65)
	fmt.Println(result2)
	result2 = case2(85)
	fmt.Println(result2)
}
