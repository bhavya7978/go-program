package main

import "fmt"

func main() {

	slice := []string{"Hello", "Penny", "Lenord"}
	slice1 := []string{"World", "Most", "Tech"}

	fmt.Println(slice)
	fmt.Println(slice1)

	final := [][]string{slice, slice1}
	fmt.Println(final)

	

}
