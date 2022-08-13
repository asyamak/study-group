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
	a.Insert(3)
	a.Insert(7)
	a.Insert(4)
	fmt.Println("before delition:")
	a.PrintInOrder()
	a.delete(5)
	fmt.Println("deletion ")
	fmt.Println("after delition:")
	a.PrintInOrder()
}

func (n *node) delete(value int) *node {
	if n == nil {
		return nil
	}
	if value > n.val {
		n.right = n.right.delete(value)
		return n
	}
	if value < n.val {
		n.left = n.left.delete(value)
		return n
	}
	if n.right == nil && n.left == nil {
		n = nil
		return nil
	}
	if n.left == nil {
		n = n.right
		return n
	}
	if n.right == nil {
		n = n.left
		return n
	}
	smallest := n.right
	for {
		if smallest != nil && smallest.left != nil {
			smallest = smallest.left
		} else {
			break
		}
	}
	n.val = smallest.val
	n.right = n.right.delete(n.val)
	return n
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
