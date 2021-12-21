package LinkList

import "fmt"

//定义结点中数据的类型为接口类型，可收任意类型数据
type Object interface {}

//定义结点的结构体
type Node struct {
	Data Object
	Next *Node
}

//定义链表的结构体
type List struct {
	HeadNode *Node
}

const MaxLength = 5

//判断链表是否为空
func (list *List) IsEmpty() bool {
	if list.HeadNode == nil {
		return true
	}

	return false
}

//获取链表的长度
func (list *List) Length() int {
	if list.HeadNode == nil {
		return 0
	}

	len := 0

	currNode := list.HeadNode
	for currNode!=nil {
		len++
		currNode = currNode.Next
	}

	return len
}

//头插法增加结点
func (list *List) AddFromHead(data Object) *Node {
	newNode := &Node{Data: data}
	if list.IsEmpty() {
		list.HeadNode = newNode
		return newNode
	}

	newNode.Next = list.HeadNode
	list.HeadNode = newNode

	return newNode
}

//尾插法增加结点
func (list *List) AddFromTail(data Object) *Node {
	newNode := &Node{Data: data}
	if list.IsEmpty() {
		list.HeadNode = newNode
		return newNode
	}

	currNode := list.HeadNode
	for currNode.Next != nil {
		currNode = currNode.Next
	}

	currNode.Next = newNode

	return newNode
}

//删除链表中指定位置的结点
func (list *List) DelNode(position int) *Node {
	if position < 0 {
		return nil
	}
	if list.IsEmpty() {
		return nil
	}
	if position > list.Length() {
		return nil
	}

	currNode := list.HeadNode

	//考虑链表只有一个结点的情况
	if currNode.Next == nil {
		list.HeadNode = nil
		return currNode
	}

	//删除头结点
	if position == 1 {
		list.HeadNode = currNode.Next
		return currNode
	}

	//删除的是尾节点
	if position == list.Length() {
		for currNode.Next.Next != nil {
			currNode = currNode.Next
		}

		currNode.Next = nil

		return currNode.Next
	}

	//删除链表中间某个结点
	preNode := list.HeadNode
	count := 1
	for count != position-1 {
		count++
		preNode = preNode.Next
	}

	preNode.Next = preNode.Next.Next

	return preNode.Next
}

//判断某个值是否在链表中
func (list *List) IsSearched(data Object) int {
	if list.IsEmpty() {
		return 0
	}

	currNode := list.HeadNode
	position := 0
	for currNode != nil {
		position++
		if currNode.Data == data {
			return position
		}
		currNode = currNode.Next
	}

	return 0
}

//遍历链表
func (list List) Traverse() {
	if list.IsEmpty() {
		fmt.Println("链表使空的")
	}

	currNode := list.HeadNode
	for currNode != nil {
		fmt.Printf("%v\t", currNode.Data)
		currNode = currNode.Next
	}
}

//链表判满
func (list List) IsFull() bool {
	if list.Length() >= 5 {
		return true
	}

	return false
}

