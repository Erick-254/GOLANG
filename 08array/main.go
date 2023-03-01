package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "peach"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Fruit list is:", len(fruitList))

	var vegetableList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegetatble list is:", vegetableList)
	fmt.Println("Vegetatble list is:", len(vegetableList))

}
