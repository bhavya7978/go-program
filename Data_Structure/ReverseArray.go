package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Enter size of your array: ")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element:\t", i)
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Your array is:", arr)
	fmt.Println(reverseArray(arr))
}

//Reverse the array
func reverseArray(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
