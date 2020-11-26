//每个目录下只能有一个package
package tree

//类似工厂函数，返回一个对象的指针（引用）
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
