package main

import (
	"data-structure/Queue/linkListQueue"
	"fmt"
)

func main() {
	queue := linkListQueue.Queue{}
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")
	queue.Enqueue("d")
	queue.Enqueue("e")
	queue.Traverse()
	queue.Dequeue()
	queue.Dequeue()
	fmt.Println()
	queue.Traverse()
}
