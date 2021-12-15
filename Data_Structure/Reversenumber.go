package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the number:")
	var input int
	fmt.Scanln(&input)

	fmt.Println("your number is:", input)
	fmt.Println(reverseNumber(input))
}

func reverseNumber(num int) int {
	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}
