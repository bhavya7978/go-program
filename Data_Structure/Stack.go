package main

import "fmt"

type Stack []string

//Check if stack is empty
func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

//Push the value in stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

//Pop-Remove the last element pushed and return the value of the stack
func (s *Stack) Pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func main() {
	var stack Stack
	fmt.Println(stack)

	stack = append(stack, "WORLD!!")
	stack = append(stack, "HELLO")
	stack = append(stack, "HI!!")

	fmt.Println(stack)

	stack.Pop()
	fmt.Println(stack)

}
