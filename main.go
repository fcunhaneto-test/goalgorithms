package main

import (
	"fmt"
	"goalgorithms/binarytree"
)

func main() {
	// a := []int{8, 2, 5, 7, 4, 6, 3, 1}
	// a := []int{25, 13, 2, 4, 7, 17, 20, 8, 15}
	// a := []int{15, 13, 2, 25, 7, 17, 20, 8, 4}
	// a := []int{-8, 4, 67, 90, 13, 54, 21}
	// a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// a := []int{0, 13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	// a := []int{2, 1, 3, 4, 1, 2, 1, 5, 4}

	var num int
	var root, current binarytree.BT
	a := []int{6, 18, 3, 7, 17, 20, 2, 4, 13, 9}

	root = binarytree.BtInit(15)

	for i := 0; i < len(a); i++ {
		root.BtInsert(a[i])
	}

	root.BtPrintAll()
	fmt.Println("***********************************")
	root.BtPrintLeft()
	fmt.Println("***********************************")
	root.BtPrintRight()
	fmt.Println("***********************************")

	fmt.Print("Enter a number to find: ")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err)
	}

	current, err = root.BtSearchNode(num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Find:", current)

	current = current.BtNodeSuccessor()
	fmt.Println("Successor:", current)

	fmt.Print("Enter a number to find: ")
	_, err = fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err)
	}

	current, err = root.BtSearchNode(num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Find:", current)

	current = current.BtNodePredecessor()
	fmt.Println("Predecessor:", current)

	current = root.BtMinimum()
	fmt.Println("Minimum:", current.Val)
	current = root.BtMaximum()
	fmt.Println("Maximum:", current.Val)
}
