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

	// var start, current exlinklist.Person
	// start.ID = 0
	// start.Next = &current
	// start.Previ = nil
	// current.ID = 0
	// current.Next = nil
	// current.Previ = &start
	// current.Next = nil

	// current = exlinklist.InsertP(1, start)
	// current = exlinklist.InsertP(2, current)
	// current = exlinklist.InsertP(3, current)
	// current = exlinklist.InsertP(4, current)
	// exlinklist.PrintP(current)

	var num int
	var root binarytree.BT
	// fmt.Println("Enter the root:")
	// fmt.Scanf("%d", &num)
	root = binarytree.InitBT(11)
	// fmt.Print("Enter the binary tree values:")
	// _, err := fmt.Scanf("%d", &num)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	a := []int{6, 19, 4, 8, 17, 43, 49, 10, 31, 5}
	for i := 0; i < len(a); i++ {
		binarytree.InsertBT(a[i], root)
		// fmt.Print("Enter the binary tree value (0 to exit):")
		// _, err := fmt.Scanf("%d", &num)
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}
	binarytree.PrintBT(root)
	fmt.Println("***********************************")
	binarytree.PrintBTLeft(root)
	fmt.Println("***********************************")
	binarytree.PrintBTRight(root)
	fmt.Println("***********************************")
	// fmt.Print("Enter a number to find:")
	// _, err := fmt.Scanf("%d", &num)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// num = binarytree.SearchInBT(num, root)
	// fmt.Println("Find:", num)
	num = binarytree.MinimumBT(root)
	fmt.Println("Minimum:", num)
	num = binarytree.MaximumBT(root)
	fmt.Println("Maximum:", num)
}
