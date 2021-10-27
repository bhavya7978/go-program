package main

import "fmt"

func main() {
	array := [5]int{}
	array[0] = 100
	array[1] = 200
	array[2] = 300
	array[3] = 400
	array[4] = 500
	for i, v := range array {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", array)
}
