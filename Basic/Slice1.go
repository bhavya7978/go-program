package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", slice)
	fmt.Println("\n", slice[:5])
	fmt.Println("\n", slice[5:])
	fmt.Println("\n", slice[2:5])
	fmt.Println("\n", slice[6:9])

	slice1 := append(slice, 101, 102, 103)
	fmt.Println(slice1)
	sliceX := []int{1111, 1112, 3333}
	slice2 := append(slice, sliceX...)
	fmt.Println(slice2)

	//DELETE=APPEND+SLICING
	fmt.Println("******************************")
	sliceOld := []int{29, 30, 31, 32, 33, 90, 65, 34, 35, 36}
	fmt.Println(sliceOld)
	sliceNew := append(sliceOld[:5], sliceOld[7:]...)
	fmt.Println(sliceNew)

	//PRINTING ALPHABETS
	fmt.Println("******************************")
	alpha := make([]string, 26, 26)
	alpha = []string{`A`, `B`, `C`, `D`, `E`, `F`, `G`, `H`, `I`, `J`, `K`, `L`, `M`, `N`, `O`, `P`, `Q`, `R`, `S`, `T`, `U`, `V`, `W`, `X`, `Y`, `Z`}

	for i := 0; i < len(alpha); i++ {
		fmt.Println(i, alpha[i])
	}
}
