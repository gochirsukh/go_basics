package main

import "fmt"

func main() {

	var age int

	fmt.Println("Enter your age:")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("You can participte in election since your age is: ", age)
	} else {
		fmt.Println("You are under age, go home")
	}
}

// https://zetcode.com/golang/scan/
