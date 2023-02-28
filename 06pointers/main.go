package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")
	// var oneptr *int
	// fmt.Println("Value of pointer is", oneptr)

	myNumber := 23

	var ptr = &myNumber //& refers the number/ pointer

	fmt.Println("Value of actual pointer is", ptr)  // memory location
	fmt.Println("Value of actual pointer is", *ptr) // actual value

	*ptr = *ptr * 2
	fmt.Println("New value is:", myNumber)

}
