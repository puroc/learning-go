package main

import "fmt"

import "pud/learning-go/tree"

func main() {
	var root tree.Node
	fmt.Println(root)

	//以下是集中创建对象的方式
	//通过 类型名{}的方式创建对象，在大括号中对对象的属性进行赋值
	root = tree.Node{Value: 1}
	//用treeNode{}创建一个对象，因为left是一个指针，所以将treeNode的地址赋值给他
	root.Left = &tree.Node{}
	root.Left.Right = &tree.Node{5, nil, nil}
	//还可以通过new(类型名)的方式来创建对象
	root.Right = new(tree.Node)
	//创建一个对象类型的切片
	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
		{0, tree.CreateNode(2), nil},
	}
	fmt.Println(nodes)

	fmt.Println("相当于java中的this")
	root.Print()

	root.SetValue(5)
	root.Print()

}
