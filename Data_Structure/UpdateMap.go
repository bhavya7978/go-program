package main

import (
	"fmt"
	"strings"
)

func main() {
	m := map[string]string{
		"First_Name": "BHAVYA",
		"Last_Name":  "jaiswal",
	}

	//Printing out the map
	fmt.Println("Before Updating:", m)
	//Convert the  last name into upper case:
	m1 := m["Last_Name"]
	map_update := strings.ToUpper(m1)
	fmt.Println("After Updating:", map_update)

	m["Last_Name"] = map_update
	fmt.Println("After Updating:", m)

}
