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
	a.delete(4)
	fmt.Println("deletion ")
	fmt.Println("after delition:")
	a.PrintInOrder()
}

func (n *node) delete(value int) *node {
	if n == nil { // if node is empty
		fmt.Println()
		return nil // do nothing
	}
	if value > n.val { // if the value we want delete is larger than current node's value
		n.right = n.right.delete(value) // move to right branch
		return n
	}
	if value < n.val { // if the value we want to delete is smaller than current node's value
		n.left = n.left.delete(value) // move to the lift branch
		return n
	}
	if n.right == nil && n.left == nil { // if left and right branch are empty
		n = nil // just delete its
		return nil
	}
	if n.left == nil { // if left node is empty
		n = n.right // move to the right node
		return n
	}
	if n.right == nil { // if right node is empty
		n = n.left // move to the left node
		return n
	}
	smallest := n.right // assign right node to variable to find it in right subtree
	for {               // infinite loop
		if smallest != nil && smallest.left != nil { // if right node is not empty and its left node is not empty
			smallest = smallest.left // change right node's to its left node value
		} else {
			break // in any other variant stop the loop
		}
	}
	n.val = smallest.val            // node's value now equal = to the smallest val in the right subtree
	n.right = n.right.delete(n.val) // and we trasfer its right node to the delete function
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
