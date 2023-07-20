package Function

import (
	"fmt"
)

type Node struct {
	Data interface{} //数据域:存放数据,interface{}空接口,可以接收任何类型
	Next *Node       //指针域:存放下一个节点地址
}
type SingleLinkedList struct {
	Head *Node //头指针:指向头结点
	Tail *Node //尾指针:指向尾节点
}

// 一: 判断链表是否为空
func (L *SingleLinkedList) IsEmpty() bool {
	return L.Head == nil
}

// 二: 遍历链表
func (L *SingleLinkedList) Traverse() {
	node := L.Head
	for ; node != nil; node = node.Next {
		fmt.Print(node.Data, " -> ")
	}
}

// 获取链表长度
func (L *SingleLinkedList) GetLength() int {
	//如果链表为空,返回0
	HeadNode := L.Head
	if HeadNode == nil {
		return 0
	}
	length := 0
	//当前节点不为nil,则length++
	for ; HeadNode != nil; HeadNode = HeadNode.Next {
		length++
	}
	return length
}

// 三: 插入节点
// 3.1尾部插入
func (L *SingleLinkedList) Append(Data interface{}) {
	node := &Node{Data: Data}
	//如果为空,直接将头指针和尾指针指向新节点
	if L.IsEmpty() {
		L.Head = node
		L.Tail = node
		return
	} else {
		//L.Tail.Next = node
		//L.Tail = node
		HeadNode := L.Head
		//从头结点开始遍历
		//插入的判断和遍历的判断有区别
		//不同之处在于插入要判断下一个节点是否为nil
		for ; HeadNode.Next != nil; HeadNode = HeadNode.Next {
		}
		HeadNode.Next = node
		L.Tail = node
	}
}

// 3.2头部插入
func (L *SingleLinkedList) InsertByHead(Data interface{}) {
	node := &Node{Data: Data}
	if L.IsEmpty() {
		L.Head = node
		L.Tail = node
	} else {
		//Go语言交换变量语法糖
		L.Head, node = node, L.Head
		//改变next指针指向
		L.Head.Next = node
		//HeadNode := L.Head
		//L.Head = node
		//L.Head.Next = HeadNode
	}
}

// InsertByIndex 中间插入
func (L *SingleLinkedList) InsertByIndex(Data interface{}, position int) {
	//如果position<=0,头部添加
	if position <= 0 {
		L.InsertByHead(Data)
	} else if position >= L.GetLength() { //position>=链表元素数量,尾部添加
		L.Append(Data)
	} else {
		node := &Node{Data: Data}
		preNode := L.Head
		count := 0
		for count != position-1 {
			preNode = preNode.Next
			count++
		}
		//顺便不能变
		//先把上一个节点的next赋值给新节点
		node.Next = preNode.Next
		//再把新节点赋给上一个节点的next
		preNode.Next = node
	}
}

// 删除节点
// 尾部删除
func (L *SingleLinkedList) DeleteHead() {
	if L.IsEmpty() {
		return
	}
	L.Head = L.Head.Next
}

// 头部删除
func (L *SingleLinkedList) DeleteTail() any {
	if L.IsEmpty() {
		return nil
	}
	HeadNode := L.Head

	//若下下个节点为nil.则下一个节点为最后一个节点
	for HeadNode.Next.Next != nil {
		HeadNode = HeadNode.Next
	}
	//将下一个节点中的数据返回
	data := HeadNode.Next.Data
	HeadNode.Next = nil
	return data
}

// 中间删除
func (L *SingleLinkedList) DeleteIndex(position int) any {
	if L.IsEmpty() {
		return nil
	}
	//positon为0,直接头部删除
	if position == 0 {
		L.DeleteHead()
	} else if position == L.GetLength()-1 { //postion为链表元素,直接尾部删除
		L.DeleteTail()
	} else {
		PreNode := L.Head
		count := 0
		for count != position-1 {
			PreNode = PreNode.Next
			count++
		}
		data := PreNode.Next.Data
		PreNode.Next = PreNode.Next.Next
		return data
	}
	return nil
}

func (L *SingleLinkedList) FindFirstCommonNode(List1 *SingleLinkedList) *Node {
	//使用Map
	Map := make(map[any]*Node)
	HeadNode := L.Head
	//遍历第一个链表,将Data和next存到Map里面
	for HeadNode != nil {
		Map[HeadNode.Data] = HeadNode.Next
		HeadNode = HeadNode.Next
	}
	//测试是否遍历正确
	fmt.Println(Map)

	//测试
	HeadNode1 := List1.Head
	//将链表中和节点挨个和Map里面地址或者值比较
	for HeadNode1 != nil {
		_, b := Map[HeadNode1.Data]
		if b {
			return HeadNode1
		}
		HeadNode1 = HeadNode1.Next
	}
	return nil
}
