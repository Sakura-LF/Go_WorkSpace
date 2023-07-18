package main

import (
	"Go_WorkSpace/Algorithm/SingleLinkedList/Function"
	"fmt"
)

func main() {
	singleLinkedList := Function.SingleLinkedList{}

	singleLinkedList.Append(114514)
	singleLinkedList.Append("Sakura")
	singleLinkedList.Append("LF")
	singleLinkedList.Append("Sony")
	singleLinkedList.InsertByHead("Java")
	singleLinkedList.InsertByIndex("Mysql", 2)

	singleLinkedList.Traverse()
	//fmt.Println(singleLinkedList.Head.Data)
	//n := singleLinkedList.GetLength()
	fmt.Println(singleLinkedList.GetLength())
	//fmt.Println(singleLinkedList.Head)
	//
	//singleLinkedList.DeleteHead()
	//
	//DeleteData := singleLinkedList.DeleteTail()
	//fmt.Println(DeleteData)
	singleLinkedList.Traverse()
	fmt.Println()
	//fmt.Println(singleLinkedList.Head)
	singleLinkedList.DeleteIndex(5)
	singleLinkedList.Traverse()

}
