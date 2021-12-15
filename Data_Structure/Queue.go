package main

import "fmt"

type Queue []string

//Check if Queue is empty
func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

//Push the value in stack
func (q *Queue) Push(str string) {
	*q = append(*q, str) // Simply append the new value to the end of the stack
}

//Pop-Remove the last element pushed and return the value of the stack
func (q *Queue) Pop() (string, bool) {
	if q.isEmpty() {
		return "", false
	} else {
		// Get the index of the top most element.
		element := (*q)[0] // Index into the slice and obtain the element.
		*q = (*q)[1:]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func main() {
	var queue Queue
	fmt.Println(queue)

	queue = append(queue, "WORLD!!")
	queue = append(queue, "HELLO")
	queue = append(queue, "HI!!")

	fmt.Println(queue)

	queue.Pop()
	fmt.Println(queue)

}
