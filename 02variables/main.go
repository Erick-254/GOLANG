package main

import "fmt"

const LoginToken string = "hgfjkd"

func main() {
	var username string = "erick"
	//variable is printed
	fmt.Println(username)
	//variable type
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = true
	//variable is printed
	fmt.Println(isLoggedIn)
	//variable type
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallValue int = 255 //methods int/uint8/16/32/64
	//variable is printed
	fmt.Println(smallValue)
	//variable type
	fmt.Printf("Variable is of type: %T\n", smallValue)

	var smallFloat float32 = 255.34543 //float64(more precision)
	//variable is printed
	fmt.Println(smallFloat)
	//variable type
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	//implicit type
	var website = "learn.io"
	fmt.Println(website)

	//no var style

	//only iside the method
	numberOfUser := 300000 //wl operator
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T\n", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)
}
