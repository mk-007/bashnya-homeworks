package main

import "fmt"

type Stack struct {
	a    [1000]int
	head int
}

func NewStack() *Stack {
	return &Stack{head: -1}
}
func (a *Stack) IsEmpty() bool {
	return a.head == -1
}
func (a *Stack) Size() int {
	return a.head + 1
}
func (a *Stack) Clear() {
	a.head = -1
}
func (a *Stack) Push(n int) {
	a.head += 1
	a.a[a.head] = n
}
func (a *Stack) Pop() {
	if a.head >= 0 {
		a.head -= 1
	}
}

func main() {
	stack := NewStack()
	for true {
		var c string
		fmt.Println("Enter command (or 'break'): ")
		fmt.Scanln(&c)
		if c == "Push" {
			fmt.Println("Enter n: ")
			var n int
			fmt.Scanln(&n)
			stack.Push(n)
			fmt.Println("Done")
		} else if c == "Pop" {
			stack.Pop()
			fmt.Println("Done")
		} else if c == "Size" {
			var n = stack.Size()
			fmt.Printf("Size = %d\n", n)
		} else if c == "IsEmpty" {
			var n = stack.IsEmpty()
			fmt.Printf("%v\n", n)
		} else if c == "Clear" {
			stack.Clear()
			fmt.Println("Done")
		} else if c == "Break" {
			break
		}
	}
}
