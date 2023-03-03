package main

import "fmt"

func main() {
	fmt.Println("Welcome to videos on slices")

	var fruitList = []string{"Apple", "Orange", "Mango"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Oange", "Banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

}
