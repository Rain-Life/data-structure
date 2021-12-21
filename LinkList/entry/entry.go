package main

import (
	"data-structure/LinkList"
	"fmt"
)

func main() {

	//验证 Lru
	Lru := LinkList.Initialize(5)
	Lru.Set("a", "1")
	Lru.Set("b", "2")
	Lru.Set("c", "3")
	Lru.Set("d", "4")
	//LinkList.Traverse()
	fmt.Println(Lru.Get("b"))



	//list1 := LinkList.List{}
	//list2 := LinkList.List{}
	//list1.AddFromTail(1)
	//list1.AddFromTail(3)
	//list1.AddFromTail(5)
	//list1.AddFromTail(7)
	//list1.AddFromTail(9)
	//list2.AddFromTail(2)
	//list2.AddFromTail(4)
	//list2.AddFromTail(6)
	//list2.AddFromTail(8)

	//if LinkList.IsPalindromeString(&list) {
	//	fmt.Println("Yes，It's palindrome string")
	//} else {
	//	fmt.Println("No, It's not palindrome string")
	//}
	//list1.Traverse()
	//fmt.Println()
	//list2.Traverse()
	//fmt.Println()
	//LinkList.MergeOrderedLinkList(&list1, &list2)
	//fmt.Println(LinkList.DelNNode(&list1, 1))
	//LinkList.ReverseLinkList(&list)
	//list.Traverse()
	//fmt.Println(list.Length())

	//fmt.Println(list.IsSearched(1))
	//fmt.Println(list.IsSearched(2))
	//fmt.Println(list.IsSearched(3))
	//fmt.Println(list.IsSearched(4))
	//fmt.Println(list.IsSearched(5))
	//list.Traverse()

	//fmt.Println(LinkList.LRUAlgorithm(&list, 6))
	//list.Traverse()
}
