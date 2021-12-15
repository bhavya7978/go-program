package main

import "fmt"

func main() {
	m := map[string][]string{
		"Corden_James":   []string{"IceCream", "Makeup"},
		"Maggie_Deliyah": []string{"Chocolate", "Chips"},
	}
	//fmt.Println(m)
	//Add record in map
	m["Styles_Harry"] = []string{"Louis", "Music", "Banana"}
	//Delete record from map
	delete(m, "Maggie_Deliyah")

	for k, v := range m {
		fmt.Println("this record is for ", k)
		for i, v2 := range v {
			fmt.Println(i, "\t", v2)
		}
	}
}
