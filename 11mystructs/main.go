package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in go//no super/parent

	erick := User{"Erick", "erick@go.dev", true, 26}
	fmt.Println(erick)
	fmt.Printf("erick details are: %+v\n", erick)
	fmt.Printf("Name is %v and email is %v.", erick.Name, erick.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
