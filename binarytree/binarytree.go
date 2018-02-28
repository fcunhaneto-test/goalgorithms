package binarytree

import (
	"errors"
	"fmt"
)

// BinaryTree a binary tree representation of integers
type BinaryTree struct {
	Val    int
	left   *BinaryTree
	right  *BinaryTree
	parent *BinaryTree
}

// BT alias for binary tree struct
type BT = *BinaryTree

// BtInit starts the binary tree by placing the root element
func BtInit(v int) BT {
	var root BinaryTree

	root.Val = v
	root.left = nil
	root.right = nil
	root.parent = nil

	return &root
}

// BtInsertNode insert a Value in binary tree
func (current BT) BtInsertNode(v int) {
	var bt BinaryTree
	bt.Val = v
	bt.left = nil
	bt.right = nil

	for true {
		if bt.Val < current.Val {
			if current.left != nil {
				current = current.left
			} else {
				bt.parent = current
				current.left = &bt
				break
			}
		} else {
			if current.right != nil {
				current = current.right
			} else {
				bt.parent = current
				current.right = &bt
				break
			}
		}
	}
}

// BtDeleteNode delete current node for binary tree
func (current BT) BtDeleteNode() {
	if current.left == nil {
		btTransplant(current, current.right)
	} else if current.right == nil {
		btTransplant(current, current.left)
	} else {
		bt := current.right.BtMinimum()
		if bt.parent != current {
			btTransplant(bt, bt.right)
			bt.right = current.right
			bt.right.parent = bt
		}
		btTransplant(current, bt)
		bt.left = current.left
		bt.left.parent = bt
	}
}

func btTransplant(u, v BT) {
	if u.parent == nil {
		u.parent = v
	} else if u == u.parent.left {
		u.parent.left = v

	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}

// BtSearchNode search a Value in binary tree
func (current BT) BtSearchNode(num int) (BT, error) {
	for current != nil {
		if num == current.Val {
			return current, nil
		} else if num < current.Val {
			current = current.left
		} else {
			current = current.right
		}
	}
	err := errors.New("Not found")
	return nil, err
}

// BtMinimum return currentmun Value in binary tree
func (current BT) BtMinimum() BT {
	for current.left != nil {
		current = current.left
	}

	return current
}

// BtMaximum return currentmun Value in binary tree
func (current BT) BtMaximum() BT {
	for current.right != nil {
		current = current.right
	}

	return current
}

// BtNodeSuccessor successor of current node
func (current BT) BtNodeSuccessor() BT {
	var bt BT
	if current.right != nil {
		current = current.right
		for current.left != nil {
			current = current.left
		}

		return current
	}

	bt = current.parent
	for bt != nil && (current == bt.right) {
		current = bt
		bt = bt.parent
	}

	return bt
}

// BtNodePredecessor successor of current node
func (current BT) BtNodePredecessor() BT {
	var bt BT
	if current.left != nil {
		current = current.left
		for current.right != nil {
			current = current.right
		}

		return current
	}

	bt = current.parent
	for bt != nil && (current == bt.left) {
		current = bt
		bt = bt.parent
	}

	return bt
}

// BtPrintAll  print all binary tree
func (current BT) BtPrintAll() {
	printBT(current)
}

// BtPrintLeft print left side of binary tree
func (current BT) BtPrintLeft() {
	current = current.left
	printBT(current)
}

// BtPrintRight  print right side of binary tree
func (current BT) BtPrintRight() {
	current = current.right
	printBT(current)
}

func printBT(current BT) {
	if current != nil {
		printBT(current.left)
		fmt.Printf("%v - %p\n", *current, current)
		printBT(current.right)
	}
}
