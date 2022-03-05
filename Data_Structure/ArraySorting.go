package main

import "fmt"

func main() {
	a := []int{10, 20, 70, 5, 90, 4}
	fmt.Println(len(a))

	for i := 0; i <= len(a)-1; i++ {
		for j := i + 1; j <= len(a)-1; j++ {

			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}

		}
	}
	fmt.Println(a)
}
