package main

import "fmt"

//该函数接收长度为5的数组，参数传递方式为值传递
func printArr(arr [5]int) {
	arr[0] = 100
	for i := range arr {
		fmt.Println(arr[i])
	}
}

//该函数接收一个长度为5的数组地址，参数传递方式相当于引用传递。
func printArr2(arr *[5]int) {
	//*arr代表arr指针所指向的变量，即main中的arr1数组。此处也可以简写为arr[0] = 100，go语言一样可以正确执行。
	(*arr)[0] = 100
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func main() {

	//使用var定义数组，可以不对数组赋初值，go会自动赋予初值
	var arr1 [5]int
	//使用:=定义的数组，需要对数组赋初值
	arr2 := [3]int{4, 5, 6}
	//如果想让go语言根据初值的数量自动指定数组长度的话，可以通过...来实现
	arr3 := [...]int{1, 2, 3}
	//定义4行5列的数组
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//遍历数组
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	//遍历数组，i代表数组的index，v代表数组中的值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	//printArr函数接收的是长度为5的数组，这里传入的arr2数组长度是3，会报错cannot use arr2 (type [3]int) as type [5]int in argument to printArr
	//printArr(arr2)

	fmt.Println("值传递。。")
	//数组是值类型参数传递，是将arr1的数据拷贝一份传递到函数中，因为printArr函数中会将数组的第一个元素修改为100，所以此处输出的arr1的第一个元素是100
	printArr(arr1)

	//数组是值类型参数传递，虽然printArr内部将传递参数arr1的第一个元素变为100了，但是main函数中的arr1值并没有改变，所以此处输出的arr1的第一个元素仍然是0
	fmt.Println(arr1)

	fmt.Println("引用传递。。")
	//此处传递的数组arr1的地址，相当于引用传递，那么在printArr2函数中修改了arr1的第一个元素，即同时会影响到main函数中的arr1
	printArr2(&arr1)
	fmt.Println(arr1)

}
