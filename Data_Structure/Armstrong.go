package main

import "fmt"

func main() {
	var input int
	result := 0
	fmt.Println("ENTER ANY 3 DIGIT NUMBER:")
	fmt.Scan(&input)
	tempNum := input

	for {
		remainder := tempNum % 10
		result += remainder * remainder * remainder
		tempNum /= 10

		if tempNum == 0 {
			break
		}
	}

	if result == input {
		fmt.Println("AMSTRONG NUMBER")
	} else {
		fmt.Println("NOT AMSTRONG NUMBER")
	}

}
