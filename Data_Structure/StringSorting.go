package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	//write a program to input a string and sort the characters into ascending order
	//sample input : xzy HkjI acb
	//output: abcHIjkxyz
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println("OUTPUT:", input)

	fmt.Println(SortStringByCharacter(input))

}

/*Soting the string
	  Convert string to []rune.
      Define a new type ByRune, which is actually []rune. Implement Len, Swap, and Less methods of type ByRune.
      Call sort. Sort to sort the slice of runes.
      Convert []rune back to string and return the string.*/
func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
