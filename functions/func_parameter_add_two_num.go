package main

import "fmt"

func main() {

	addtwo(1, 3)
	firstLastName("John", "Wick")
}

func addtwo(num1, num2 int) {
	sum := num1 + num2
	fmt.Println("The sum of two is", sum)
}

func firstLastName(first, last string) {
	fmt.Println("Name of the actor is", first, last)
}
