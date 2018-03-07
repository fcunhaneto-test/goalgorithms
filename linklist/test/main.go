package main

import (
	"fmt"
	"goalgorithms/linklist"
)

func main() {
	var p linklist.Person
	var head, tail, current linklist.LL
	var head1, tail1 linklist.LL

	p.ID = 1
	p.Name = "Francisco"
	head = linklist.LlInit(p)

	p.ID = 2
	p.Name = "Erica"
	tail = linklist.LlInsertTail(p, head)

	p.ID = 3
	p.Name = "João"
	tail = linklist.LlInsertTail(p, tail)

	p.ID = 4
	p.Name = "Paula"
	tail = linklist.LlInsertTail(p, tail)

	current = linklist.LlFindFromHead(3, head)

	// linklist.DeleteNode(current)

	fmt.Println()
	linklist.LlPrint(head)
	fmt.Println()
	linklist.LlPrintReverse(tail)
	fmt.Println()

	p.ID = 5
	p.Name = "Rachel"
	head1 = linklist.LlInit(p)

	p.ID = 6
	p.Name = "Tercio"
	tail1 = linklist.LlInsertTail(p, head1)

	linklist.LlInsertLlAfter(head1, tail1, current)

	fmt.Println()
	linklist.LlPrint(head)
	fmt.Println()
	linklist.LlPrintReverse(tail)
	fmt.Println()

	current = linklist.LlFindFromHead(2, head)

	linklist.LlDeleteNode(current)

	fmt.Println()
	linklist.LlPrint(head)
	fmt.Println()
	linklist.LlPrintReverse(tail)
	fmt.Println()

	// fmt.Println()
	// for current.Next != nil {
	// 	current = linklist.LlGetNext(current)
	// 	p = current.V.(Person)
	// 	fmt.Println(p.ID)
	// 	fmt.Println(p.Name)
	// }

	// var num int
	// var root, current binarytree.BT
	// a := []int{6, 18, 3, 7, 17, 20, 2, 4, 13, 9}

	// root = binarytree.BtInit(15)
	// root.BtInsertNode(678)
	// for i := 0; i < len(a); i++ {
	// 	root.BtInsertNode(a[i])
	// }

	// // root.BtPrintAll()
	// // fmt.Println("***********************************")
	// // root.BtPrintLeft()
	// // fmt.Println("***********************************")
	// // root.BtPrintRight()
	// // fmt.Println("***********************************")

	// fmt.Print("Enter a number to find: ")
	// _, err := fmt.Scanf("%d", &num)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// current, err = root.BtSearchNode(num)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Find:", current)

	// current.BtDeleteNode()

	// root.BtPrintLeft()
	// // current = current.BtMinimum()
	// // fmt.Println("Minimun:", current)

	// fmt.Print("Enter a number to find: ")
	// _, err = fmt.Scanf("%d", &num)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// current, err = root.BtSearchNode(num)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Find:", current)

	// current = current.BtNodePredecessor()
	// fmt.Println("Predecessor:", current)

	// current = root.BtMinimum()
	// fmt.Println("Minimum:", current.Val)
	// current = root.BtMaximum()
	// fmt.Println("Maximum:", current.Val)

	// fmt.Println("***********************************")
	// root.BtPrintAll()
	// fmt.Println("***********************************")
	// root.BtPrintLeft()
	// fmt.Println("***********************************")
	// root.BtPrintRight()
	// fmt.Println("***********************************")

	// var n int
	// fmt.Println("Entre com um numero:")
	// fmt.Scanf("%d", &n)
	// fmt.Println(n)
}
