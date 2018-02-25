package binarytree

import (
	"fmt"
)

// BinaryTree a binary tree representation of integers
type BinaryTree struct {
	val   int
	left  *BinaryTree
	right *BinaryTree
}

// BT alias for binary tree struct
type BT = *BinaryTree

// InitBT starts the binary tree by placing the root element
func InitBT(v int) BT {
	var root BinaryTree

	root.val = v
	root.left = nil
	root.right = nil

	return &root
}

//InsertBT insert a value in binary tree
func InsertBT(v int, current BT) {
	var bt BinaryTree
	bt.val = v
	bt.left = nil
	bt.right = nil

	for true {
		if bt.val < current.val {
			if current.left != nil {
				current = current.left
			} else {
				current.left = &bt
				break
			}
		} else {
			if current.right != nil {
				current = current.right
			} else {
				current.right = &bt
				break
			}
		}
	}
}

// SearchInBT search a value in binary tree
func SearchInBT(num int, current BT) int {
	for current != nil {
		if num == current.val {
			return current.val
		} else if num < current.val {
			current = current.left
		} else {
			current = current.right
		}
	}

	return 0
}

// MinimumBT return currentmun value in binary tree
func MinimumBT(current BT) int {
	var num int
	for current != nil {
		num = current.val
		current = current.left
	}

	return num
}

// MaximumBT return currentmun value in binary tree
func MaximumBT(current BT) int {
	var num int
	for current != nil {
		num = current.val
		current = current.right
	}

	return num
}

// PrintBT  print all binary tree
func PrintBT(current BT) {
	if current != nil {
		fmt.Println(current)
		PrintBT(current.left)
		PrintBT(current.right)
	}
}

// PrintBTLeft print left side of binary tree
func PrintBTLeft(current BT) {
	fmt.Println(current)
	current = current.left
	if current != nil {
		fmt.Println(current)
		PrintBT(current.left)
		PrintBT(current.right)
	}
}

// PrintBTRight  print right side of binary tree
func PrintBTRight(current BT) {
	fmt.Println(current)
	current = current.right
	if current != nil {
		fmt.Println(current)
		PrintBT(current.left)
		PrintBT(current.right)
	}
}
