//deadlock channel
package main

import "fmt"

//OUTPUT:all goroutines are asleep - deadlock!
func main() {
	//Solution1:using anonymous func
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	/*Buffer Channel
	c := make(chan int, 1)
	//Solution1:using anonymous func
	 c <- 42
	fmt.Println(<-c)
	*/
	//***********************************************************//
	//Exercise2

	//cs := make(chan<- int)---Output:cannot receive from send-only channel cs (variable of type chan<- int)
	cs := make(chan int) //----Make this bidirectional channel
	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

}
