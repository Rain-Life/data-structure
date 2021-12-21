package LinkList

import "fmt"

//单链表实现LRU算法
func LRUAlgorithm(list *List, data interface{}) Object {
	position := list.IsSearched(data)
	if position > 0 {//1. 数据在链表中
		//删除该结点
		list.DelNode(position)
		//将该节点插入到头部
		list.AddFromHead(data)

		return data
	}

	//2. 数据不在链表中
	if list.IsFull() { //链表满了
		//删除最后一个结点
		list.DelNode(list.Length())

		//并将当前结点插入到头部'
		list.AddFromHead(data)
	} else {//链表没满
		//直接插入头部即可
		list.AddFromHead(data)
	}

	return data
}

// 通过链表来存储一个字符串，判断这个字符串是不是回文字符串（Palindrome String）
func IsPalindromeString(list *List) bool {
	if list.IsEmpty() {
		return false
	}
	if list.Length() == 1 {
		return true
	}

	fast := list.HeadNode
	slow := list.HeadNode
	next := list.HeadNode
	newListNode := &Node{}
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next = slow.Next
		slow.Next = newListNode
		newListNode = slow
		slow = next
	}

	if fast != nil { // 比如有三个结点的时候
		slow = slow.Next
	}

	for slow != nil {
		if slow.Data != newListNode.Data {
			return false
		}

		newListNode = newListNode.Next
		slow = slow.Next
	}

	return true
}


// 单链表反转（Reverse）
func ReverseLinkList(list *List) {
	if list.IsEmpty() {
		fmt.Println("链表是空的")
		return
	}
	if list.Length() == 1 {
		fmt.Println("链表只有一个结点")
		return
	}

	//这里使用哨兵的方式来实现
	newHead := &Node{}
	newHead.Next = list.HeadNode
	prevNode := newHead.Next
	currNode := prevNode.Next
	for currNode != nil {
		prevNode.Next = currNode.Next
		currNode.Next = newHead.Next
		newHead.Next = currNode
		currNode = prevNode.Next
	}
	list.HeadNode = newHead.Next
}


// 链表中环的检测（Check Ring）
func CheckRing(list *List) bool {
	if list.IsEmpty() {
		return false
	}

	if list.Length() == 1 && list.HeadNode.Next == list.HeadNode {
		return true
	}

	fast := list.HeadNode
	slow := list.HeadNode
	for fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == nil {
			return false
		}
		if fast == slow {
			return true
		}
	}

	return false
}


// 两个有序的链表合并（Merge Order LinkList）
func MergeOrderedLinkList(list1, list2 *List) {
	if list1.HeadNode == nil && list2.HeadNode == nil {
		fmt.Println("俩链表都是空的")
		return
	}

	newList := &List{}
	if list1.HeadNode == nil && list2.HeadNode != nil {
		newList.HeadNode = list2.HeadNode
		newList.Traverse()
		return
	}
	if list1.HeadNode != nil && list2.HeadNode == nil {
		newList.HeadNode = list1.HeadNode
		newList.Traverse()
		return
	}

	currNode1 := list1.HeadNode
	currNode2 := list2.HeadNode
	if currNode1.Data.(int) < currNode2.Data.(int) {
		newList.HeadNode = currNode1
		currNode1 = currNode1.Next
	} else {
		newList.HeadNode = currNode2
		currNode2 = currNode2.Next
	}

	newHead := newList.HeadNode
	currNodeNew := newHead

	for currNode1 != nil && currNode2 != nil {
		if currNode1.Data.(int) < currNode2.Data.(int) {
			currNodeNew.Next = currNode1
			currNode1 = currNode1.Next
		} else {
			currNodeNew.Next = currNode2
			currNode2 = currNode2.Next
		}

		currNodeNew = currNodeNew.Next
	}

	if currNode1 == nil {
		currNodeNew.Next = currNode2
	}
	if currNode2 == nil {
		currNodeNew = currNode1
	}

	newList.Traverse()
}



// 删除单向链表倒数第n个结点（Delete Node）
func DelNNode(list *List, lastN int) Object {
	if list.IsEmpty() {
		fmt.Println("链表是空的")
		return nil
	}

	length := list.Length()
	if lastN <0 || lastN > length {
		fmt.Println("参数不合法")
		return nil
	}

	position := length - lastN + 1
	return list.DelNode(position)
}



// 求链表的中间结点（Find Middle Node）
func MiddleNode(list *List) Object {
	//还是用快慢指针的方式
	if list.IsEmpty() {
		fmt.Println("链表是空的")
		return 0
	}
	if list.Length() == 1 {
		return list.HeadNode.Data
	}

	fast := list.HeadNode
	slow := list.HeadNode
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow.Data
}


// 约瑟夫问题（Joseph Circle）
func JosephCircle(list *List, number int) Object {
	if list.IsEmpty() {
		fmt.Println("链表是空的")
		return nil
	}
	if list.Length() == 1 {
		return list.HeadNode.Data
	}

	currNode := list.HeadNode
	preNode := list.HeadNode
	count := 1
	for preNode != list.HeadNode {
		preNode = preNode.Next
	}
	for currNode.Next != list.HeadNode {
		if count == number {
			preNode.Next = currNode.Next
			currNode = preNode.Next
			count = 1
		} else {
			count++
		}

		preNode = currNode
		currNode = currNode.Next
	}

	return currNode.Data
}



