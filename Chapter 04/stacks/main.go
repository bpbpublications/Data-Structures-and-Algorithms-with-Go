package main

import (
	"fmt"
	"stack/stack"
)

func main() {
	var s stack.Stack

	s.Push(27)
	s.Push(5)
	s.Push(1)
	s.Push(18)

	fmt.Println(s)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	fmt.Println(s)
}
