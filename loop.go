package main

import "fmt"

func sum(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		result += i
	}
	return result
}

//for后面什么都不加，就是死循环
func deadLoop() {
	//for {
	//
	//}
}

func main() {
	fmt.Println(sum(5))
}
