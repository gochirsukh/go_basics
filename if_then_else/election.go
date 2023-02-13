package main

import "fmt"

func main() {

	fmt.Println("Enter your age:")
	var age int
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("Who are you voting for?: ")
		var candidate string
		fmt.Scanln(&candidate)
		if candidate == "Trump" {
			fmt.Println("You are republican")
		} else {
			fmt.Println("You are democrat")
		}
	} else {
		fmt.Printf("You are %v years old. You are too young to vote", age)
	}

}
