package main

import (
	"Go_WorkSpace/Algorithm/SingleLinkedList/Function"
	"fmt"
)

func main() {
	singleLinkedList1 := Function.SingleLinkedList{}
	singleLinkedList2 := Function.SingleLinkedList{}

	data1 := "A1"
	data2 := "A2"
	data3 := "C1"
	data4 := "C2"
	data5 := "C3"

	data6 := "B1"
	data7 := "B2"
	data8 := "B3"

	//node1 := Function.Node{Data: "A1"}
	//node2 := Function.Node{Data: "A2"}
	//node3 := Function.Node{Data: "C1"}
	//node4 := Function.Node{Data: "C2"}
	//node5 := Function.Node{Data: "C3"}
	//node6 := Function.Node{Data: "B1"}
	//node7 := Function.Node{Data: "B2"}
	//node8 := Function.Node{Data: "B3"}

	//Function.Node{Data: 1}

	//A链表
	singleLinkedList1.Append(data1)
	singleLinkedList1.Append(data2)
	singleLinkedList1.Append(data3)
	singleLinkedList1.Append(data4)
	singleLinkedList1.Append(data5)

	//B链表&
	singleLinkedList2.Append(data6)
	singleLinkedList2.Append(data7)
	singleLinkedList2.Append(data8)
	singleLinkedList2.Append(data3)
	singleLinkedList2.Append(data4)
	singleLinkedList2.Append(data5)

	singleLinkedList1.Traverse()
	fmt.Println("\n")
	singleLinkedList2.Traverse()
	fmt.Println("\n")

	fmt.Println(singleLinkedList1.FindFirstCommonNode(&singleLinkedList2))

}
