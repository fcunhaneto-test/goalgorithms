package main

import (
	"fmt"
	"goalgorithms/linklist"
)

func main() {
	var s string
	var p linklist.Person
	var head, tail, current linklist.LL

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

	fmt.Println()
	linklist.LlPrint()
	fmt.Println()
	linklist.LlPrintReverse()
	fmt.Println()
	fmt.Scanf("%s", &s)

	current = linklist.LlFind(3)
	fmt.Println("Find 3:", current)

	p.ID = 5
	p.Name = "Rachel"
	current = linklist.LlInsertBefore(p, current)
	// fmt.Printf("%v - %p\n", current.Previ, current.Previ)
	// fmt.Printf("%v - %p\n", current, current)
	// fmt.Printf("%v - %p\n", current.Next, current.Next)
	fmt.Println()
	linklist.LlPrint()
	fmt.Println()
	linklist.LlPrintReverse()
	fmt.Scanf("%s", &s)

	p.ID = 6
	p.Name = "Tercio"
	head = linklist.LlInsertBefore(p, head)

	fmt.Println()
	linklist.LlPrint()
	fmt.Println()
	linklist.LlPrintReverse()
	fmt.Scanf("%s", &s)

	current = linklist.LlFind(2)
	fmt.Println("Find 2:", current)

	linklist.LlDeleteNode(current)

	fmt.Println()
	linklist.LlPrint()
	fmt.Println()
	linklist.LlPrintReverse()
	fmt.Scanf("%s", &s)

	linklist.LlDeleteNode(head)
	head = linklist.GetHead()
	fmt.Println("Head:", head)

	fmt.Println()
	linklist.LlPrint()
	fmt.Println()
	linklist.LlPrintReverse()
	fmt.Scanf("%s", &s)

	linklist.LlDeleteNode(tail)

	fmt.Println()
	linklist.LlPrint()
	fmt.Println()
	linklist.LlPrintReverse()
	fmt.Println()

	tail = linklist.GetTail()
	fmt.Println("Tail:", tail)

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
