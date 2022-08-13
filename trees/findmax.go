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
	c := a.FindMax()
	fmt.Println("maximum: ", c)
	a.PrintInOrder()
}

func (n *node) FindMax() int {
	if n.right == nil {
		return n.val
	}
	return n.right.FindMax()
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
