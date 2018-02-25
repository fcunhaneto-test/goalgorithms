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
func InsertBT(v int, atual BT) {
	var bt BinaryTree
	bt.val = v
	bt.left = nil
	bt.right = nil

	for true {
		if bt.val < atual.val {
			if atual.left != nil {
				atual = atual.left
			} else {
				atual.left = &bt
				break
			}
		} else {
			if atual.right != nil {
				atual = atual.right
			} else {
				atual.right = &bt
				break
			}
		}
	}
}

// SearchInBT search a value in binary tree
func SearchInBT(num int, atual BT) BT {
	for atual != nil {
		if num == atual.val {
			return atual
		} else if num < atual.val {
			atual = atual.left
		} else {
			atual = atual.right
		}
	}

	return nil
}

// PrintBT  print all binary tree
func PrintBT(atual BT) {
	if atual != nil {
		fmt.Println(atual)
		PrintBT(atual.left)
		PrintBT(atual.right)
	}
}

// PrintBTLeft print left side of binary tree
func PrintBTLeft(atual BT) {
	atual = atual.left
	if atual != nil {
		fmt.Println(atual)
		PrintBT(atual.left)
		PrintBT(atual.right)
	}
}

// PrintBTRight  print right side of binary tree
func PrintBTRight(atual BT) {
	atual = atual.right
	if atual != nil {
		fmt.Println(atual)
		PrintBT(atual.left)
		PrintBT(atual.right)
	}
}
