package main

import "fmt"

func main() {

	m := map[string]int{
		"Louis": 29,
		"Harry": 27,
	}
	fmt.Println(m) //Printing the map
	fmt.Println(m["Louis"])

	//Comma OK idioms to check particular value exists in map or not
	if v, ok := m["Liam"]; ok {
		fmt.Println(v)
	}

	if v, ok := m["Harry"]; ok {
		fmt.Println(v)
	}
}
