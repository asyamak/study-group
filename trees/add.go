package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func main() {
	a := &node{val: 5}
	a.Insert(10)
	a.Insert(12)
	// a.Insert(5)
	a.Insert(1)
	a.Insert(11)
	a.Insert(2)
	a.Insert(6)
	a.PrintInOrder()
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

func (n *node) PrintInOrder() {
	if n == nil {
		return
	}
	n.left.PrintInOrder()
	fmt.Println(n.val)
	n.right.PrintInOrder()
}
