package main

import "fmt"

//当参数中的数组类型不写长度时，相当于是一个切片，切片本身是没有数据的，他相当于是底层数据的一个引用。
//所以在函数中修改切片中数据的值，会影响到main函数中再次访问切片获得数据的结果
func updateSlice(arr []int) {
	arr[0] = 100
}

//切片中有por、len、cap三个属性，其中por是指针，len是切片的长度，cap是切片的容量，cap不足时会自动乘2.
func main() {
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//输出数组[2,6)区间的元素
	fmt.Println(arr1[2:6])
	//输出数组[0,6)区间的元素
	fmt.Println(arr1[:6])
	//输出数组[2,最后一个元素]区间的元素
	fmt.Println(arr1[2:])
	//输出数组[0,最后一个元素]区间的元素
	fmt.Println(arr1[:])

	//修改切片中的数据
	updateSlice(arr1[:])
	fmt.Println(arr1[:])

	fmt.Println("向arr1中添加元素")
	//向切片添加元素时，会返回一个新的切片，不会影响原来切片中的数据
	arr2 := append(arr1[:], 10)
	arr3 := append(arr2[:], 11)
	fmt.Println(arr1[:])
	fmt.Println(arr2[:])
	fmt.Println(arr3[:])

	//创建切片的几种方式
	//创建切片，并赋予初值
	s1 := []int{2, 4, 6, 8}
	//创建长度为16的切片，但不赋予初值，由go语言自动赋予初值
	s2 := make([]int, 16)
	//创建长度为16，容量为32的切片，但不赋予初值，由go语言自动赋予初值
	s3 := make([]int, 16, 32)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println("拷贝元素")
	//把s1的内容拷贝到s2中
	copy(s2, s1)
	fmt.Println(s2)

	fmt.Println("删除元素")
	//删除中间元素，即将被删除元素的前后两部分切片合并在一起
	s5 := append(s2[:3], s2[4:]...)
	fmt.Println(s5)

	//删除头元素
	front := s1[0]
	s1 = s1[1:]

	//删除尾元素
	tail := s1[len(s1)-1]
	s1 = s1[:len(s1)-1]
	fmt.Println(front, tail, s1)

}
