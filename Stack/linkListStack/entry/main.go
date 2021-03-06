package main

import (
	"data-structure/Stack/linkListStack"
	"fmt"
)

func main()  {
	stack := linkListStack.Stack{}
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Push("d")
	stack.Push("e")

	stack.Traverse()
	fmt.Println()
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop())
}
