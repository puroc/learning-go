package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//这里的(node Node)叫接收者，相当于调用java中的this，即调用print函数的那个对象。但这里的接收者是值传递，所以在这个方法中修改参数中的值是不会影响main函数中的变量的。
func (node Node) Print() {
	fmt.Println(node.Value)
}

//在go语言中只有值传递，没有引用传递。那么想要在函数中设置参数对象的值，并影响main函数中的变量，则可以通过指针来实现，效果就相当于java中的引用传递。
func (node *Node) SetValue(value int) {
	node.Value = value
}
