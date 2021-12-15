package main

import "fmt"

func main() {
	slice := []int{-5, 10, 8, 9, 1, 12, 100}
	maxNum, minNum := maxAndMin(slice)
	fmt.Println(maxNum, minNum)

}
func maxAndMin(slice []int) (int, int) {

	max := slice[0]
	min := slice[0]

	for _, v := range slice {
		if v < min {
			min = v
		}
	}

	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max, min
}
