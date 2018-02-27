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

// InitBt starts the binary tree by placing the root element
func InitBt(v int) BT {
	var root BinaryTree

	root.Val = v
	root.left = nil
	root.right = nil
	root.parent = nil

	return &root
}

// InsertBt insert a Value in binary tree
func (current BT) InsertBt(v int) {
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

// SearchInBt search a Value in binary tree
func (current BT) SearchInBt(num int) (BT, error) {
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

// MinimumBt return currentmun Value in binary tree
func (current BT) MinimumBt() BT {
	for current.left != nil {
		current = current.left
	}

	return current
}

// MaximumBt return currentmun Value in binary tree
func (current BT) MaximumBt() BT {
	for current.right != nil {
		current = current.right
	}

	return current
}

// SuccessorBt successor of current node
func (current BT) SuccessorBt() BT {
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

// PredecessorBt successor of current node
func (current BT) PredecessorBt() BT {
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

// PrintBtAll  print all binary tree
func (current BT) PrintBtAll() {
	printBT(current)
}

// PrintBtLeft print left side of binary tree
func (current BT) PrintBtLeft() {
	current = current.left
	printBT(current)
}

// PrintBtRight  print right side of binary tree
func (current BT) PrintBtRight() {
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
