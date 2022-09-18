package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func main() {
	a := &node{val: 13}
	a.Insert(10)
	a.Insert(7)
	a.Insert(11)
	a.Insert(6)
	a.Insert(18)
	a.Insert(16)
	a.Insert(20)
	a.Insert(22)
	a.Insert(19)
	a.Insert(17)

	a.print()
	a.remove(7)
	a.print()
}

func (n *node) Insert(value int) {
	if n == nil {
		fmt.Println("tree is empty")
	}
	if value == n.val {
		fmt.Println("value already exist")
	}
	if value < n.val {
		if n.left == nil {
			n.left = &node{val: value}
		}
		n.left.Insert(value)
	}

	if value > n.val {
		if n.right == nil {
			n.right = &node{val: value}
		}
		n.right.Insert(value)
	}
}

func (n *node) print() {
	if n == nil {
		return
	}
	n.left.print()
	fmt.Println(n.val)
	n.right.print()
}

func (n *node) delete(value int) {
	n.remove(value)
}

func (n *node) remove(value int) *node {
	if n == nil {
		return nil
	}
	if value < n.val {
		n.left = n.left.remove(value)
		return n
	}
	if value > n.val {
		n.right = n.right.remove(value)
		return n
	}
	if n.left == nil && n.right == nil {
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
	smallestValOnRight := n.right
	for {
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}
	n.val = smallestValOnRight.val
	n.right = n.right.remove(n.val)
	return n
}
