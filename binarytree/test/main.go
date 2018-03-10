package main

import (
	"fmt"
	"goalgorithms/binarytree"
)

func main() {
	var num int
	var root, current binarytree.BT
	a := []int{6, 18, 3, 7, 17, 20, 2, 4, 13, 9}

	root = binarytree.BtInit(15)
	root.BtInsertNode(678)
	for i := 0; i < len(a); i++ {
		root.BtInsertNode(a[i])
	}

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

	current.BtDeleteNode()

	root.BtPrintLeft()
	// current = current.BtMinimum()
	// fmt.Println("Minimun:", current)

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

	fmt.Println("***********************************")
	root.BtPrintAll()
	fmt.Println("***********************************")
	root.BtPrintLeft()
	fmt.Println("***********************************")
	root.BtPrintRight()
	fmt.Println("***********************************")
}
