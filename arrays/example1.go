package main

import "fmt"

func main() {

	arrayOfIntegers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrayOfIntegers)
	fmt.Println("Access index", arrayOfIntegers[3])

	arrayOfStrings := [3]string{"red", "green", "black"}
	fmt.Println(arrayOfStrings)
	fmt.Println("Accessed array of strings", arrayOfStrings[2])

}
