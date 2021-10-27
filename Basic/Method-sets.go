package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello World!")

}

func saySomething(h human) {
	h.speak()

}

func main() {
	p1 := person{
		first: "Robin",
	}
	/*saySomething(p1) //doesnt work:cannot use p1 (variable of type person) as human value in argument
	to saySomething: missing method speak (speak has pointer receiver)*/
	saySomething(&p1)
	//p1.speak().......//Any variable of type person have access to speak()
}
