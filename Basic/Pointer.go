package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "Cary"
	//(*p).name = "Blueberry"
}
func main() {

	p1 := person{
		name: "James",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
