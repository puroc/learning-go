package main

import (
	"fmt"
	"reflect"
	"runtime"
)

//函数返回多个返回值
func div(a, b int) (int, int) {
	return a / b, a % b
}

//给返回值起个名字，起不起名字对于函数的调用者来说没有区别
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return q, r
}

//在go语言中的函数通常返回两个返回值，第一个返回值是结果，第二个返回值是error，如果执行时没有错误时，error为nil
func eval(a, b int, op string) (int, error) {
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
		return 0, fmt.Errorf("unsupport operator:" + op)
	}
	return result, nil
}

func apply(op func(a, b int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	name := runtime.FuncForPC(p).Name()
	fmt.Printf("calling func name:%s ,param:(%d,%d)\n", name, a, b)
	return op(a, b)
}

func add(a, b int) int {
	return a + b
}

//可变参数列表
func sum2(numbers ...int) int {
	result := 0
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

func main() {
	fmt.Println(div(13, 3))
	q, r := div2(13, 3)
	fmt.Println(q, r)
	//div2函数有两个返回值，其中第二个返回值我不想用，但是go语言要求声明过的变量必须要被使用，否则就编译失败，这时可以通过下划线"_"来占位，代替不想使用的返回值。如下所示
	a, _ := div2(13, 3)
	fmt.Println(a)
	if result, err := eval(1, 2, "&"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(apply(add, 2, 3))
	fmt.Println(sum2(1, 2, 3))

}
