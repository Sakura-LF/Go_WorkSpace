package Linklist

type Node struct {
	Data     interface{} //节点存储的数据
	NextNode *Node       // 指向下一个节点
}

type LinkedList struct {
	Head *Node //头节点
	Tail *Node //尾节点
}
