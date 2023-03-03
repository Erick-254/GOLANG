package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to videos on slices")

	var fruitList = []string{"Apple", "Orange", "Mango"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Oange", "Banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 934
	highScore[1] = 334
	highScore[2] = 434
	highScore[3] = 934

	highScore = append(highScore, 555, 666, 777, 888, 999) // apppend can accomodate all values

	fmt.Println(highScore)
	//only available in slice
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

}
