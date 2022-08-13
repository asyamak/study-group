package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func main() {
	a := &node{val: 5}
	a.Insert(6)
	a.Insert(8)
	a.Insert(2)
	a.Insert(10)
	a.Insert(4)
	c := a.FindMin()
	fmt.Println("minimum: ", c)
	a.PrintInOrder()
}

func (n *node) FindMin() int {
	if n.left == nil {
		return n.val
	}
	return n.left.FindMin()
}

func (n *node) PrintInOrder() {
	if n == nil {
		return
	}
	n.left.PrintInOrder()
	fmt.Println(n.val)
	n.right.PrintInOrder()
}

func (n *node) Insert(value int) {
	if n == nil {
		fmt.Println("Tree is empty")
		// return
	}

	if value == n.val {
		fmt.Println("value already exist")
		// return
	}
	if value > n.val {
		if n.right == nil {
			n.right = &node{val: value}
			return
		}
		n.right.Insert(value)

	}
	if value < n.val {
		if n.left == nil {
			n.left = &node{val: value}
			return
		}
		n.left.Insert(value)

	}
}
