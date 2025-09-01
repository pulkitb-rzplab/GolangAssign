package main

import "fmt"

type Node struct {
	value string
	left  *Node
	right *Node
}

func newnode(val string) *Node {
	new := Node{value: val, left: nil, right: nil}
	return &new
}

func exampleTree() *Node {
	//for a+b*d-c
	a := newnode("A")
	b := newnode("B")
	c := newnode("C")
	d := newnode("D")

	plus := newnode("+")
	minus := newnode("-")
	multiply := newnode("*")

	minus.left = plus
	minus.right = c
	plus.left = a
	plus.right = multiply
	multiply.left = b
	multiply.right = d

	return minus

}

func preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf(root.value + " ")
	preorder(root.left)
	preorder(root.right)

}

func postorder(root *Node) {
	if root == nil {
		return
	}
	postorder(root.left)
	postorder(root.right)
	fmt.Printf(root.value + " ")

}

//func main() {

	root := exampleTree()
	preorder(root)
	fmt.Print("\n\n")
	postorder(root)
}
